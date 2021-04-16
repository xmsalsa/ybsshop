/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 15:03
 */
package merchant

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"reflect"
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	merchant2 "shop/application/models/merchant"
	"shop/application/service/merchant"
)

func GetMerchantlst(ctx iris.Context) {
	page := _package.Page{
		Page:  ctx.URLParamIntDefault("page", _package.PAGE),
		Limit: ctx.URLParamIntDefault("limit", _package.LIMIT),
	}
	where := make(map[string]interface{})
	Auth := _package.GetAuthority(ctx)
	slic := make([]interface{}, 0)
	where["pid"] = append(slic, "pid", "=", Auth.MerchId)
	where["is_flag"] = append(slic, "is_flag", "=", 1)
	mer := new(merchant.MerchantService)
	List := mer.GetMerchantlst(where, _package.Sql(ctx), page)
	if mer.Error() != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, List, mer.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, List, response.NoErr.Msg))
	}
	return
}

/** 添加子商户 **/
func PostCreateSon(ctx iris.Context) {
	Authority := _package.GetAuthority(ctx)
	param := merchant.PostMerchantCreate{
		Id: 0,
	}
	ctx.ReadJSON(&param)
	param.Pid = Authority.MerchId
	param.IsFlag = 1
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, _package.StrErrs(err, reflect.TypeOf(param), "error")))
			return
		}
	}
	mer := new(merchant.MerchantService)
	data := merchant2.Merchant{}
	easygorm.Begin()
	if param.Id == 0 {
		if param.MerPassword == "" {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, "请输入商户密码"))
			return
		}
		data = mer.PostCreateSon(param)
	} else {
		update := _package.GetUpdateInit(ctx)
		data = mer.PostUpdateSon(update, param)
	}
	if mer.Error() != "" {
		easygorm.Rollback()
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, mer.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
		easygorm.Commit()
	}
	return
}

func RegisteredMerchant(ctx iris.Context) {
	param := merchant.RegisteredMerchant{}
	ctx.ReadJSON(&param)
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, _package.StrErrs(err, reflect.TypeOf(param), "error")))
			return
		}
	}
	mer := new(merchant.MerchantService)
	easygorm.Begin()
	data := mer.RegisteredMerchant(param)
	if mer.Error() != "" {
		easygorm.Rollback()
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, mer.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
		easygorm.Commit()
	}
	return
}
