/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 10:02:28
 */
package controllers

import (
	"shop/application/libs/response"
	"shop/application/service/user"

	"github.com/kataras/iris/v12"
)

// 调试专用: 遇神杀神 佛挡杀佛 魔来斩魔__Bug
func Debug(ctx iris.Context) {

	// 测试用户
	param := user.UserDetail{
		Phone: ctx.URLParamDefault("phone", ""),
		Uid:   ctx.URLParamIntDefault("uid", 0),
	}

	u, err := user.GetUser(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(u))
	return
}
