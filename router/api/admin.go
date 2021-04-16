/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 00:25:27
 */
package api

import (
	"shop/application/controllers"
	cadmin "shop/application/controllers/admin"

	"github.com/kataras/iris/v12"
)

func AdminRouter(admin iris.Party) {
	admin.PartyFunc("/account", func(account iris.Party) {
		account.Get("/exist", controllers.AccountExist).Name = "判断账号是否存在"
	})
	admin.Post("/create", cadmin.AdminCreate).Name = "管理员创建"
}
