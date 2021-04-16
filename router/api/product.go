/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 14:32
 */
package api

import (
	"github.com/kataras/iris/v12"
	"shop/application/controllers/product"
)

func ProductRouter(productroter iris.Party) {
	productroter.PartyFunc("/rule", func(rule iris.Party) {
		rule.Get("/", product.GetRule).Name = "获取商品规格列表"
		rule.Post("/", product.PostRule).Name = "添加商品规格"
		rule.Delete("/", product.DeleteRule).Name = "删除商品规格"
	})
	productroter.PartyFunc("/category", func(rule iris.Party) {
		rule.Get("/", product.GetCategory).Name = "获取商品分类列表"
		rule.Post("/", product.PostCategor).Name = "添加商品分类"
		rule.Delete("/{id:int}", product.DeleteCategory).Name = "删除商品分类"
	})
	productroter.PartyFunc("/product", func(rule iris.Party) {
		rule.Post("/", product.PostProduct).Name = "添加商品"

		rule.Get("/", product.GetProduct).Name = "获取商品列表"
		rule.Get("/details", product.GetDetails).Name = "获取商品详情"

		rule.Put("/set_show", product.PutSetShow).Name = "上下架商品"
		rule.Put("/product_show", product.PutUnShow).Name = "批量上下架商品"

	})
	return
}
