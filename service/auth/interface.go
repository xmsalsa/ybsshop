package auth

import (
	"errors"
	"time"
)

const (
	ZxwSessionTokenPrefix           = "ZST:"
	ZxwSessionBindAccountPrefix     = "ZSBU:"
	ZxwSessionAccountPrefix         = "ZSU:"
	ZxwSessionAccountMaxTokenPrefix = "ZXWAccountMaxToken"
)

var (
	ErrTokenInvalid                  = errors.New("token is invalid")
	ZxwSessionAccountMaxTokenDefault = 10
)

const (
	NoneScope uint64 = iota
	AdminScope
)

const (
	NoAuth int = iota
	AuthPwd
	AuthCode
	AuthThirdParty
)

const (
	LoginTypeWeb int = iota
	LoginTypeApp
	LoginTypeWx
	LoginTypeAlipay
	LoginApplet
)

var (
	RedisSessionTimeoutWeb    = 30 * time.Minute
	RedisSessionTimeoutApp    = 24 * time.Hour
	RedisSessionTimeoutWx     = 5 * 52 * 168 * time.Hour
	RedisSessionTimeoutApplet = 7 * 24 * time.Hour
)

type Session struct {
	AccountId    string `json:"account_id" redis:"account_id"`
	LoginType    int    `json:"login_type" redis:"login_type"`
	AuthType     int    `json:"auth_type" redis:"auth_type"`
	CreationDate int64  `json:"creation_data" redis:"creation_data"`
	ExpiresIn    int    `json:"expires_in" redis:"expires_in"`
	Scope        uint64 `json:"scope" redis:"scope"`
}

// Authentication  认证
type Authentication interface {
	ToCache(token string, id uint64) error
	SyncAccountTokenCache(token string) error
	DelAccountTokenCache(token string) error
	AccountTokenExpired(token string) error
	UpdateAccountTokenCacheExpire(token string) error
	GetSessionV2(token string) (*Session, error)
	GetAuthId(token string) (uint, error)
	IsAccountTokenOver(token string) bool
	CleanAccountTokenCache(token string) error
	Close()
}
