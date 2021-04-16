/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 21:36:54
 */
package client

import (
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	sclient "shop/application/service/client"

	"github.com/kataras/iris/v12"
)

func SaveSetGroup(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := sclient.SSaveSetGroup{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	err = sclient.SaveSetGroup(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(make(map[string]int)))
	return
}
