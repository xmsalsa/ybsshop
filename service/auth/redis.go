package auth

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"shop/application/libs"
	"shop/application/libs/logging"
	"shop/service/cache"
	"strconv"
	"strings"
	"time"
)

type RedisAuth struct {
	Conn *cache.RedisCluster
}

func NewRedisAuth() *RedisAuth {
	return &RedisAuth{
		Conn: cache.GetRedisClusterClient(),
	}
}

// GetAuthId
func (ra *RedisAuth) GetAuthId(token string) (uint, error) {
	sess, err := ra.GetSessionV2(token)
	if err != nil {
		return 0, err
	}
	id := uint(libs.ParseInt(sess.AccountId, 10))
	return id, nil
}

//  GetSessionV2 session
func (ra *RedisAuth) GetSessionV2(token string) (*Session, error) {
	sKey := ZxwSessionTokenPrefix + token
	if !ra.Conn.Exists(sKey) {
		return nil, ErrTokenInvalid
	}
	pp := new(Session)
	if err := ra.Conn.LoadRedisHashToStruct(sKey, pp); err != nil {
		return nil, err
	}
	return pp, nil
}

// IsAccountTokenOver 超过登录设备限制
func (ra *RedisAuth) IsAccountTokenOver(accountId string) bool {
	logging.DebugLogger.Debugf("account token count ", ra.getAccountTokenCount(accountId), " account max count ", ra.getAccountTokenMaxCount())
	if ra.getAccountTokenCount(accountId) >= ra.getAccountTokenMaxCount() {
		return true
	}
	return false
}

// getAccountTokenCount 获取登录数量
func (ra *RedisAuth) getAccountTokenCount(accountId string) int {
	count, err := redis.Int(ra.Conn.Scard(ZxwSessionAccountPrefix + accountId))
	if err != nil {
		logging.ErrorLogger.Errorf("get account token count err: %+v", err)
		return 0
	}
	return count
}

// getAccountTokenMaxCount 最大登录限制
func (ra *RedisAuth) getAccountTokenMaxCount() int {
	count, err := redis.Int(ra.Conn.GetKey(ZxwSessionAccountMaxTokenPrefix))
	if err != nil {
		return ZxwSessionAccountMaxTokenDefault
	}
	return count
}

// AccountTokenExpired 过期 token
func (ra *RedisAuth) AccountTokenExpired(token string) error {
	uKey := ZxwSessionBindAccountPrefix + token
	sKeys, err := redis.Strings(ra.Conn.Members(uKey))
	if err != nil {
		logging.ErrorLogger.Errorf("account token expired get members err: %+v", err)
		return err
	}
	for _, v := range sKeys {
		if !strings.Contains(v, ZxwSessionAccountPrefix) {
			continue
		}
		_, err = ra.Conn.Do("SREM", v, token)
		if err != nil {
			logging.ErrorLogger.Errorf("account token expired do srem err: %+v", err)
			return err
		}
	}
	if _, err = ra.Conn.Del(uKey); err != nil {
		logging.ErrorLogger.Errorf("account token expired del err: %+v", err)
		return err
	}
	return nil
}

// getAccountScope 角色
func GetAccountScope(accountType string) uint64 {
	switch accountType {
	case "admin":
		return AdminScope
	}
	return NoneScope
}

// ToCache 缓存 token
func (ra *RedisAuth) ToCache(token string, id uint64) error {
	sKey := ZxwSessionTokenPrefix + token
	rsv2 := &Session{
		AccountId:       strconv.FormatUint(id, 10),
		LoginType:    LoginTypeWeb,
		AuthType:     AuthPwd,
		CreationDate: time.Now().Unix(),
		Scope:        GetAccountScope("admin"),
	}
	if _, err := ra.Conn.HMSet(sKey,
		"account_id", rsv2.AccountId,
		"login_type", rsv2.LoginType,
		"auth_type", rsv2.AuthType,
		"creation_data", rsv2.CreationDate,
		"expires_in", rsv2.ExpiresIn,
		"scope", rsv2.Scope,
	); err != nil {
		logging.ErrorLogger.Errorf("to cache token err: %+v", err)
		return err
	}

	return nil
}

// SyncAccountTokenCache 同步 token 到缓存
func (ra *RedisAuth) SyncAccountTokenCache(token string) error {
	rsv2, err := ra.GetSessionV2(token)
	if err != nil {
		return err
	}
	sKey := ZxwSessionAccountPrefix + rsv2.AccountId
	if _, err := ra.Conn.Sadd(sKey, token); err != nil {
		logging.ErrorLogger.Errorf("sync account token cache sadd err: %+v", err)
		return err
	}
	sKey2 := ZxwSessionBindAccountPrefix + token
	_, err = ra.Conn.Sadd(sKey2, sKey)
	if err != nil {
		logging.ErrorLogger.Errorf("sync account token cache sadd err: %+v", err)
		return err
	}
	return nil
}

//UpdateAccountTokenCacheExpire 更新过期时间
func (ra *RedisAuth) UpdateAccountTokenCacheExpire(token string) error {
	rsv2, err := ra.GetSessionV2(token)
	if err != nil {
		return err
	}
	if rsv2 == nil {
		return errors.New("token cache is nil")
	}
	if _, err = ra.Conn.Expire(ZxwSessionTokenPrefix+token, int(ra.getTokenExpire(rsv2).Seconds())); err != nil {
		logging.ErrorLogger.Errorf("update account token cache expire err: %+v", err)
		return err
	}
	return nil
}

// getTokenExpire 过期时间
func (ra *RedisAuth) getTokenExpire(rsv2 *Session) time.Duration {
	timeout := RedisSessionTimeoutApp
	if rsv2.LoginType == LoginTypeWeb {
		timeout = RedisSessionTimeoutWeb
	} else if rsv2.LoginType == LoginTypeWx {
		timeout = RedisSessionTimeoutWx
	} else if rsv2.LoginType == LoginTypeAlipay {
		timeout = RedisSessionTimeoutWx
	}
	return timeout
}

// DelAccountTokenCache 删除token缓存
func (ra *RedisAuth) DelAccountTokenCache(token string) error {
	rsv2, err := ra.GetSessionV2(token)
	if err != nil {
		return err
	}
	if rsv2 == nil {
		return errors.New("token cache is nil")
	}
	sKey := ZxwSessionAccountPrefix + rsv2.AccountId
	_, err = ra.Conn.Do("SREM", sKey, token)
	if err != nil {
		logging.ErrorLogger.Errorf("del account token cache do srem err: %+v", err)
		return err
	}
	err = ra.DelTokenCache(token)
	if err != nil {
		return err
	}

	return nil
}

// DelTokenCache 删除token缓存
func (ra *RedisAuth) DelTokenCache(token string) error {
	sKey2 := ZxwSessionBindAccountPrefix + token
	_, err := ra.Conn.Del(sKey2)
	if err != nil {
		logging.ErrorLogger.Errorf("del token cache del key err: %+v", err)
		return err
	}

	sKey3 := ZxwSessionTokenPrefix + token
	_, err = ra.Conn.Del(sKey3)
	if err != nil {
		logging.ErrorLogger.Errorf("del token cache del key err: %+v", err)
		return err
	}

	return nil
}

// CleanAccountTokenCache 清空token缓存
func (ra *RedisAuth) CleanAccountTokenCache(token string) error {
	rsv2, err := ra.GetSessionV2(token)
	if err != nil {
		logging.ErrorLogger.Errorf("clean account token cache member err: %+v", err)
		return err
	}
	sKey := ZxwSessionAccountPrefix + rsv2.AccountId
	var allTokens []string
	allTokens, err = redis.Strings(ra.Conn.Members(sKey))
	if err != nil {
		logging.ErrorLogger.Errorf("clean account token cache member err: %+v", err)
		return err
	}
	_, err = ra.Conn.Del(sKey)
	if err != nil {
		logging.ErrorLogger.Errorf("clean account token cache del err: %+v", err)
		return err
	}

	for _, token := range allTokens {
		err = ra.DelTokenCache(token)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close
func (ra *RedisAuth) Close() {
	ra.Conn.Close()
}
