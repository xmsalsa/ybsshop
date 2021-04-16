package auth

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"shop/application/libs"
	"shop/application/libs/logging"
	"strconv"
	"strings"
	"time"
)

type tokens []string
type skeys []string

var localCache *cache.Cache

type LocalAuth struct {
	Cache *cache.Cache
}

func NewLocalAuth() *LocalAuth {
	if localCache == nil {
		localCache = cache.New(4*time.Hour, 24*time.Minute)
	}
	return &LocalAuth{
		Cache: localCache,
	}
}

// GetAuthId
func (la *LocalAuth) GetAuthId(token string) (uint, error) {
	sess, err := la.GetSessionV2(token)
	if err != nil {
		return 0, err
	}
	id := uint(libs.ParseInt(sess.AccountId, 10))
	return id, nil
}

func (la *LocalAuth) ToCache(token string, id uint64) error {
	sKey := ZxwSessionTokenPrefix + token
	rsv2 := &Session{
		AccountId:       strconv.FormatUint(id, 10),
		LoginType:    LoginTypeWeb,
		AuthType:     AuthPwd,
		CreationDate: time.Now().Unix(),
		Scope:        GetAccountScope("admin"),
	}
	la.Cache.Set(sKey, rsv2, la.getTokenExpire(rsv2))
	return nil
}

func (la *LocalAuth) SyncAccountTokenCache(token string) error {
	rsv2, err := la.GetSessionV2(token)
	if err != nil {
		logging.ErrorLogger.Errorf("SyncAccountTokenCache err: %+v\n", err)
		return err
	}

	sKey := ZxwSessionAccountPrefix + rsv2.AccountId
	ts := tokens{}
	if uTokens, uFound := la.Cache.Get(sKey); uFound {
		ts = uTokens.(tokens)
	}
	ts = append(ts, token)

	la.Cache.Set(sKey, ts, la.getTokenExpire(rsv2))

	sKey2 := ZxwSessionBindAccountPrefix + token
	sys := skeys{}
	if keys, found := la.Cache.Get(sKey2); found {
		sys = keys.(skeys)
	}
	sys = append(sys, sKey)
	la.Cache.Set(sKey2, sys, la.getTokenExpire(rsv2))
	return nil
}

func (la *LocalAuth) DelAccountTokenCache(token string) error {
	rsv2, err := la.GetSessionV2(token)
	if err != nil {
		return err
	}
	if rsv2 == nil {
		return errors.New("token cache is nil")
	}
	sKey := ZxwSessionAccountPrefix + rsv2.AccountId
	exp := la.getTokenExpire(rsv2)
	if utokens, ufound := la.Cache.Get(sKey); ufound {
		t := utokens.(tokens)
		for index, u := range t {
			if u == token {
				utokens = append(t[0:index], t[index:]...)
				la.Cache.Set(sKey, utokens, exp)
			}
		}
	}
	err = la.DelTokenCache(token)
	if err != nil {
		return err
	}

	return nil
}

// DelTokenCache 删除token缓存
func (la *LocalAuth) DelTokenCache(token string) error {
	la.Cache.Delete(ZxwSessionBindAccountPrefix + token)
	la.Cache.Delete(ZxwSessionTokenPrefix + token)
	return nil
}

func (la *LocalAuth) AccountTokenExpired(token string) error {
	rsv2, err := la.GetSessionV2(token)
	if err != nil {
		return err
	}
	if rsv2 == nil {
		return errors.New("token cache is nil")
	}

	exp := la.getTokenExpire(rsv2)
	uKey := ZxwSessionBindAccountPrefix + token
	if sKeys, found := la.Cache.Get(uKey); !found {
		return errors.New("token skey is empty")
	} else {
		for _, v := range sKeys.(skeys) {
			if !strings.Contains(v, ZxwSessionAccountPrefix) {
				continue
			}
			if utokens, ufound := la.Cache.Get(v); ufound {
				t := utokens.(tokens)
				for index, u := range t {
					if u == token {
						utokens = append(t[0:index], t[index:]...)
						la.Cache.Set(v, utokens, exp)
					}
				}
			}
		}
	}

	la.Cache.Delete(uKey)
	return nil
}

func (la *LocalAuth) UpdateAccountTokenCacheExpire(token string) error {
	rsv2, err := la.GetSessionV2(token)
	if err != nil {
		return err
	}
	if rsv2 == nil {
		return errors.New("token cache is nil")
	}
	la.Cache.Set(ZxwSessionTokenPrefix+token, rsv2, la.getTokenExpire(rsv2))

	return nil
}

// getTokenExpire 过期时间
func (la *LocalAuth) getTokenExpire(rsv2 *Session) time.Duration {
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

func (la *LocalAuth) GetSessionV2(token string) (*Session, error) {
	sKey := ZxwSessionTokenPrefix + token
	get, _ := la.Cache.Get(sKey)
	logging.DebugLogger.Infof("GetSessionV2: %+v", get)
	if food, found := la.Cache.Get(sKey); !found {
		logging.ErrorLogger.Errorf("get serssion err ", ErrTokenInvalid)
		return nil, ErrTokenInvalid
	} else {
		return food.(*Session), nil
	}
}

func (la *LocalAuth) IsAccountTokenOver(AccountId string) bool {
	logging.DebugLogger.Debugf("account token count ", la.getAccountTokenCount(AccountId), " account max count ", la.getAccountTokenMaxCount())
	if la.getAccountTokenCount(AccountId) >= la.getAccountTokenMaxCount() {
		return true
	}
	return false
}

// getAccountTokenCount 获取登录数量
func (la *LocalAuth) getAccountTokenCount(AccountId string) int {
	if accountTokens, found := la.Cache.Get(ZxwSessionAccountPrefix + AccountId); !found {
		return 0
	} else {
		return len(accountTokens.(tokens))
	}
}

// getAccountTokenMaxCount 最大登录限制
func (la *LocalAuth) getAccountTokenMaxCount() int {
	if count, found := la.Cache.Get(ZxwSessionAccountMaxTokenPrefix); !found {
		return ZxwSessionAccountMaxTokenDefault
	} else {
		return count.(int)
	}
}

// CleanAccountTokenCache 清空token缓存
func (la *LocalAuth) CleanAccountTokenCache(token string) error {
	rsv2, err := la.GetSessionV2(token)
	if err != nil {
		logging.ErrorLogger.Errorf("clean Account token cache member err: %+v", err)
		return err
	}
	sKey := ZxwSessionAccountPrefix + rsv2.AccountId
	if accountTokens, found := la.Cache.Get(sKey); !found {
		return nil
	} else {
		for _, token := range accountTokens.(tokens) {
			err = la.DelTokenCache(token)
			if err != nil {
				return err
			}
		}
	}
	la.Cache.Delete(sKey)

	return nil
}

// 兼容 redis
func (la *LocalAuth) Close() {}
