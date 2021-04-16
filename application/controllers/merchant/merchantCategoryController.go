/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 14:31
 */
package merchant

import (
	"github.com/kataras/iris/v12"
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/service/merchant"
)

func GetMerchantCategorylst(ctx iris.Context) {
	page := _package.Page{
		Page:  ctx.URLParamIntDefault("page", _package.PAGE),
		Limit: ctx.URLParamIntDefault("limit", _package.LIMIT),
	}
	mer := new(merchant.MerchantCategoryService)
	List := mer.GetMerchantCategorylst(_package.Sql(ctx), page)
	if mer.Error() != "" {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, List, mer.Error()))
	} else {
		ctx.JSON(response.NewResponse(response.NoErr.Code, List, response.NoErr.Msg))
	}
	return
}
