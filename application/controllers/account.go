package controllers

import (
	"strconv"
	"strings"
	"time"

	"shop/application/libs"
	"shop/application/libs/easygorm"
	"shop/application/libs/logging"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	"shop/application/models"
	"shop/application/service/account"
	"shop/service/dao"
	"shop/service/dao/daccount"

	"github.com/kataras/iris/v12"
)

func Profile(ctx iris.Context) {
	id, err := dao.GetAuthId(ctx)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	profile := &daccount.AccountResponse{}
	err = profile.Profile(id)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	ctx.JSON(response.NewResponse(response.NoErr.Code, profile, response.NoErr.Msg))
}

type Avatar struct {
	Avatar string
}

func ChangeAvatar(ctx iris.Context) {
	id, err := dao.GetAuthId(ctx)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	avatar := &Avatar{}
	if err := ctx.ReadJSON(avatar); err != nil {
		logging.ErrorLogger.Errorf("change avatar read json error ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	validErr := libs.Validate.Struct(*avatar)
	errs := libs.ValidRequest(validErr)
	if len(errs) > 0 {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, strings.Join(errs, ";")))
		return
	}
	err = easygorm.GetEasyGormDb().Model(&models.Account{}).Where("id = ?", id).Update("avatar", avatar.Avatar).Error
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
}

func GetAccount(ctx iris.Context) {
	info := daccount.AccountResponse{}
	err := dao.Find(&info, ctx)
	if err != nil {
		logging.ErrorLogger.Errorf("get account get err ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, info, response.NoErr.Msg))
}

func CreateAccount(ctx iris.Context) {
	accountReq := &daccount.AccountReq{}
	if err := ctx.ReadJSON(accountReq); err != nil {
		logging.ErrorLogger.Errorf("create account read json err ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	validErr := libs.Validate.Struct(*accountReq)
	errs := libs.ValidRequest(validErr)
	if len(errs) > 0 {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, strings.Join(errs, ";")))
		return
	}

	err := dao.Create(&daccount.AccountResponse{}, ctx, map[string]interface{}{
		"Name":      accountReq.Name,
		"Username":  accountReq.Username,
		"Password":  libs.HashPassword(accountReq.Password),
		"Intro":     accountReq.Intro,
		"Avatar":    accountReq.Avatar,
		"CreatedAt": time.Now(),
		"UpdatedAt": time.Now(),
	})
	if err != nil {
		logging.ErrorLogger.Errorf("create account get err ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}

	ctx.JSON(response.NewResponse(response.NoErr.Code, accountReq, response.NoErr.Msg))
	return
}

func UpdateAccount(ctx iris.Context) {
	accountReq := &daccount.AccountReq{}
	if err := ctx.ReadJSON(accountReq); err != nil {
		logging.ErrorLogger.Errorf("create account read json err ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	validErr := libs.Validate.Struct(*accountReq)
	errs := libs.ValidRequest(validErr)
	if len(errs) > 0 {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, strings.Join(errs, ";")))
		return
	}

	err := dao.Update(&daccount.AccountResponse{}, ctx, map[string]interface{}{
		"Name":      accountReq.Name,
		"Password":  libs.HashPassword(accountReq.Password),
		"Intro":     accountReq.Intro,
		"Avatar":    accountReq.Avatar,
		"UpdatedAt": time.Now(),
	})
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

func DeleteAccount(ctx iris.Context) {
	err := dao.Delete(&daccount.AccountResponse{}, ctx)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}

	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

// GetAccounts
func GetAccounts(ctx iris.Context) {
	name := ctx.FormValue("name")
	page, _ := strconv.Atoi(ctx.FormValue("page"))
	pageSize, _ := strconv.Atoi(ctx.FormValue("pageSize"))
	orderBy := ctx.FormValue("orderBy")
	sort := ctx.FormValue("sort")

	list, err := dao.All(&daccount.AccountResponse{}, ctx, name, sort, orderBy, page, pageSize)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, list, response.NoErr.Msg))
	return
}

func AccountExist(ctx iris.Context) {
	// 参数处理
	username := ctx.URLParamTrim("username")
	param := account.SAccountExsit{
		Username: username,
	}
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	exsit, err := account.AccountExist(username)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]bool{"exsit": exsit}))
	return
}
