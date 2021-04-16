/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 14:55
 */
package product

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	product2 "shop/application/models/product"
	"shop/application/service/product"
	"time"
)

func GetRule(ctx iris.Context) {
	page := _package.Page{
		Page:  ctx.URLParamIntDefault("page", _package.PAGE),
		Limit: ctx.URLParamIntDefault("limit", _package.LIMIT),
	}
	List, err := product.GetListRule(ctx.URLParamDefault("rule_name", ""), _package.Sql(ctx), page)
	if err != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, List, err))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, List, response.NoErr.Msg))
	return
}

func PostRule(ctx iris.Context) {
	user := product.RuleServiceParam{
		Id: 0,
	}
	ctx.ReadJSON(&user)
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		ctx.JSON(response.NewResponse(response.NoErr.Code, _package.EmptyDta(), "结构体参数不全"))
		return
	}
	spec, e := json.Marshal(user.Spec)
	if e != nil {
		ctx.JSON(response.NewResponse(response.NoErr.Code, _package.EmptyDta(), err.Error()))
		return
	}
	GetAuthority := _package.GetAuthority(ctx)
	Model := product2.StoreProductRule{
		Id:        user.Id,
		MerchId:   GetAuthority.MerchId,
		RuleValue: string(spec),
		RuleName:  user.RuleName,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Effect:    1,
	}
	Model, err = product.AnyAddRule(Model)
	if err != nil {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, _package.EmptyDta(), err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, Model, response.NoErr.Msg))
	return
}

/**删除sku模板*/
func DeleteRule(ctx iris.Context) {
	dele := product.AttrRuleDelete{
		All: 0,
	}
	ctx.ReadJSON(&dele)
	List, err := product.DelectRule(dele, _package.Sql(ctx))
	if err != nil {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, _package.EmptyDta(), err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, List, response.NoErr.Msg))
	return
}
