/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 00:39:53
 */
package admin

import (
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	sadmin "shop/application/service/admin"

	"github.com/kataras/iris/v12"
)

func AdminCreate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	// 参数处理  暂默认添加的是商户管理员
	param := sadmin.SAdminCreate{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
		UseApp:     0, // 不可使用APP
		UsePc:      1, // 可登录PC
		Type:       2, // 商户管理员
		Status:     1, // 正常
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	data := make(map[string]interface{})
	data, err = sadmin.AdminCreate(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(data))
	return
}
