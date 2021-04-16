/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 18:01
 */
package product

import (
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	product2 "shop/application/models/product"
	"shop/application/service/product"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

/**添加商品分类*/
func PostCategor(ctx iris.Context) {
	Authority := _package.GetAuthority(ctx)
	param := product2.StoreCategory{
		Id:        0,
		MerchId:   Authority.MerchId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Effect:    _package.EFFECT,
	}
	ctx.ReadJSON(&param)
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, "结构体参数错误"))
		return
	}
	param, err = product.AnyAddCategory(param)
	if err != nil {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, param, response.NoErr.Msg))
	return
}

/**获取商品分类*/
func GetCategory(ctx iris.Context) {
	where := make(map[string]string)
	Params := ctx.URLParams()
	Params["pid"] = ctx.URLParamDefault("pid", "0")
	for key, value := range Params {
		if value == "" {
			continue
		}
		if key == "pid" || key == "is_show" || key == "cate_name" {
			where[key] = value
		}
	}
	page := _package.Page{
		Page:  ctx.URLParamIntDefault("page", _package.PAGE),
		Limit: ctx.URLParamIntDefault("limit", _package.LIMIT),
	}
	List, err := product.GetListCategory(where, _package.Sql(ctx), page)
	if err != nil {
		ctx.JSON(response.NewResponse(response.NoErr.Code, List, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, List, response.NoErr.Msg))
	return
}

/**删除商品分类*/
func DeleteCategory(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	List, err := product.DelectCategory(id, _package.Sql(ctx), _package.GetUpdateInit(ctx))
	if err != nil {
		ctx.JSON(response.NewResponse(response.NoErr.Code, _package.EmptyDta(), err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, List, response.NoErr.Msg))
	return
}
