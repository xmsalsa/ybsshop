package controllers

import (
	"fmt"
	"shop/service/dao/daccount"
	"strings"

	"shop/application/libs"
	"shop/application/libs/easygorm"
	"shop/application/libs/logging"
	"shop/application/libs/response"
	"shop/application/models"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
)

type LoginRe struct {
	Username string `json:"username" validate:"required,gte=2,lte=50" comment:"用户名"`
	Password string `json:"password" validate:"required,gte=6,lte=30"  comment:"密码"`
}

type Account struct {
	Id       uint
	Username string
	Password string
}

type Token struct {
	AccessToken string
}

func Login(ctx iris.Context) {
	loginReq := &LoginRe{}
	if err := ctx.ReadJSON(loginReq); err != nil {
		logging.ErrorLogger.Errorf("login read request json err ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	logging.DebugLogger.Debugf("login Account ", loginReq)

	validErr := libs.Validate.Struct(*loginReq)
	errs := libs.ValidRequest(validErr)
	if len(errs) > 0 {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, strings.Join(errs, ";")))
		return
	}

	account := Account{Username: loginReq.Username}
	err := easygorm.GetEasyGormDb().Model(models.Account{}).Where("username = ?", loginReq.Username).Find(&account).Error
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}

	logging.DebugLogger.Debugf("account", account)

	if account.Id == 0 {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, nil, fmt.Sprintf("用户 %s 不存在", account.Username)))
		return
	}

	if ok := bcrypt.Match(loginReq.Password, account.Password); !ok {
		ctx.JSON(response.NewResponse(response.AuthErr.Code, nil, "用户名或密码错误"))
		return
	}

	var token string
	token, err = daccount.Login(uint64(account.Id))
	if err != nil {
		ctx.JSON(response.NewResponse(response.AuthErr.Code, nil, err.Error()))
		return
	}

	logging.DebugLogger.Debugf("account token %s", token)

	ctx.JSON(response.NewResponse(response.NoErr.Code, &Token{AccessToken: token}, response.NoErr.Msg))
	return
}

func Logout(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	err := daccount.Logout(value.Raw)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

func Expire(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	if err := daccount.Expire(value.Raw); err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

func Clear(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	if err := daccount.Clear(value.Raw); err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}
