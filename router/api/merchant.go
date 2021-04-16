/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 14:51
 */
package api

import (
	"github.com/kataras/iris/v12"
	"shop/application/controllers/merchant"
)

func MerchantRouter(merchat iris.Party) {
	merchat.PartyFunc("/", func(rule iris.Party) {
		rule.Get("/lst", merchant.GetMerchantlst).Name = "获取商户列表"
		rule.Post("/create", merchant.PostCreateSon).Name = "添加商户"
		rule.Post("/register", merchant.RegisteredMerchant).Name = "注册商户"
	})
	merchat.PartyFunc("/category", func(rule iris.Party) {
		rule.Get("/lst", merchant.GetMerchantCategorylst).Name = "获取行业分类列表"
	})
	return
}
