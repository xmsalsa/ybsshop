package daccount

import (
	"errors"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/iris-contrib/middleware/jwt"
	"shop/application/libs/logging"
	"shop/service/auth"
	"shop/service/dao"
)

// Login 登录
func Login(id uint64) (string, error) {
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	if authDriver.IsAccountTokenOver(strconv.FormatUint(id, 10)) {
		return "", errors.New("以达到同时登录设备上限")
	}
	// 使用分布唯一算法
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat": node.Generate(),
	})
	var tokenString string
	tokenString, err = token.SignedString([]byte("HS2JDFKhu7Y1av7b"))
	if err != nil {
		logging.ErrorLogger.Errorf("signed string token err", err)
		return "", err
	}
	if err = authDriver.ToCache(tokenString, id); err != nil {
		logging.ErrorLogger.Errorf("to cache account token err", err)
		return "", err
	}
	if err = authDriver.SyncAccountTokenCache(tokenString); err != nil {
		logging.ErrorLogger.Errorf("sync account token err", err)
		return "", err
	}

	err = dao.CreateOplog("认证", dao.ActionLogin, "", uint(id))
	if err != nil {
		logging.ErrorLogger.Errorf("login add oplog get err ", err)
		return "", err
	}

	return tokenString, nil
}

// Logout 退出
func Logout(token string) error {
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	id, err := authDriver.GetAuthId(token)
	if err != nil {
		logging.ErrorLogger.Errorf("logout get auth id err", err)
		return err
	}
	err = authDriver.DelAccountTokenCache(token)
	if err != nil {
		logging.ErrorLogger.Errorf("logout del account token err", err)
		return err
	}
	err = dao.CreateOplog("认证", dao.ActionLogout, "", id)
	if err != nil {
		logging.ErrorLogger.Errorf("logout add oplog get err ", err)
		return err
	}
	return nil
}

// Expire 更新
func Expire(token string) error {
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	if err := authDriver.UpdateAccountTokenCacheExpire(token); err != nil {
		logging.ErrorLogger.Errorf("update account token err", err)
		return err
	}
	return nil
}

// Check
func Check(token string) (*auth.Session, error) {
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	rsv2, err := authDriver.GetSessionV2(token)
	if err != nil {
		logging.ErrorLogger.Errorf("check account token err", err)
		return nil, err
	}
	return rsv2, nil
}

// Clear 清除
func Clear(token string) error {
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	err := authDriver.CleanAccountTokenCache(token)
	if err != nil {
		logging.ErrorLogger.Errorf("check account token err", err)
		return err
	}
	return nil
}
