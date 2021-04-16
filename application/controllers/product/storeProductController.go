/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-8
 * Time: 下午 15:27
 */
package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"reflect"
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/service/product"
)

/** 添加商品分类 **/
func PostProduct(ctx iris.Context) {
	//Authority := _package.GetAuthority(ctx)
	param := product.PostProduct{}
	ctx.ReadJSON(&param)
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, _package.StrErrs(err, reflect.TypeOf(param), "errors")))
			return
		}
	}
	pro := new(product.ProductService)
	pro.Sql = *easygorm.GetEasyGormDb().Begin()
	data := pro.CreateProduct(param, _package.GetAuthority(ctx))
	if pro.Error() != "" {
		pro.Sql.Rollback()
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, pro.Error()))
	} else {
		pro.Sql.Commit()
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
	}

	return
}

/** 获取商品分类 **/
func GetProduct(ctx iris.Context) {
	where := make(map[string]interface{})
	slic := make([]interface{}, 0)
	switch ctx.URLParamIntDefault("type", 1) {
	case 1: //出售中的商品
		where["is_show"] = append(slic, "is_show", "=", 1)
		where["stock"] = append(slic, "stock", ">", 0)
		where["effect"] = append(slic, "effect", "=", 1)
		break
	case 2: //仓库中的商品
		where["is_show"] = append(slic, "is_show", "=", 0)
		where["effect"] = append(slic, "effect", "=", 1)
	case 4: //已经售罄的商品
		where["stock"] = append(slic, "stock", "=", 0)
		where["effect"] = append(slic, "effect", "=", 1)
	case 5: //警戒库存商品
		where["stock"] = append(slic, "stock", "<=", 10)
		where["effect"] = append(slic, "effect", "=", 1)
	case 6: //回收站商品
		where["effect"] = append(slic, "effect", "=", 0)
	default:
		ctx.JSON(response.NewResponse(response.NoErr.Code, _package.EmptyDta(), "暂无此类型"))
		return
	}
	storeName := ctx.URLParamDefault("store_name", "")
	if storeName != "" {
		where["store_name"] = append(slic, "store_name", "like", "%"+storeName+"%")
	}
	page := _package.Page{
		Page:  ctx.URLParamIntDefault("page", _package.PAGE),
		Limit: ctx.URLParamIntDefault("limit", _package.LIMIT),
	}
	cateId := ctx.URLParamIntDefault("cate_id", 0)
	if cateId != 0 {
		cat, error := product.FindStoreCategory(cateId)
		if error != nil {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, cat, error.Error()))
			return
		}
		wheres := make(map[string]interface{})
		if cat.Pid != 0 {
			wheres["cate_id"] = append(slic, "cate_id", "=", cat.Id)
		} else {
			wheres["cate_id"] = append(slic, "cate_pid", "=", cat.Id)
		}
		//则查询父id分类下的商品
		pro := new(product.ProductCatesService)
		data := pro.GetProductCates(wheres, _package.Sql(ctx), page)
		wheres["id"] = append(slic, "id", "in", data)
	}
	pro := new(product.ProductService)
	data := pro.GetProduct(where, _package.Sql(ctx), page)
	if pro.Error() != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, pro.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
	}
	return
}

/** 上下架商品 **/
func PutSetShow(ctx iris.Context) {
	param := product.PutSetShow{}
	ctx.ReadJSON(&param)
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, _package.StrErrs(err, reflect.TypeOf(param), "errors")))
			return
		}
	}
	pro := new(product.ProductService)
	data := pro.PutSetShow(param.Id, param.IsShow, _package.GetAuthority(ctx))
	if pro.Error() != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, pro.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
	}
	return
}

/** 批量上下架商品 **/
func PutUnShow(ctx iris.Context) {
	param := product.PutUnShow{}
	ctx.ReadJSON(&param)
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, _package.StrErrs(err, reflect.TypeOf(param), "errors")))
			return
		}
	}
	pro := new(product.ProductService)
	data := pro.PutUnShow(param, _package.GetUpdateInit(ctx), _package.Sql(ctx))
	if pro.Error() != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, pro.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
	}
	return
}

/** 获取商品详情 **/
func GetDetails(ctx iris.Context) {
	param := product.GetDetails{}
	ctx.ReadJSON(&param)
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, _package.StrErrs(err, reflect.TypeOf(param), "errors")))
			return
		}
	}
	pro := new(product.ProductService)
	data := pro.GetDetails(param)
	if pro.Error() != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, data, pro.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, data, response.NoErr.Msg))
	}
	return
}
