/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-15
 * Time: 下午 21:42
 */
package product

import (
	"github.com/kataras/iris/v12"
)

/**	添加虚拟（回复/评论）	*/
func SaveFictitiousReply(ctx iris.Context) {
	//Authority := _package.GetAuthority(ctx)
	//param := product2.StoreCategory{
	//	Id:        0,
	//	MerchId:   Authority.MerchId,
	//	CreatedAt: time.Now().Unix(),
	//	UpdatedAt: time.Now().Unix(),
	//	Effect:    _package.EFFECT,
	//}
	//ctx.ReadJSON(&param)
	//validate := validator.New()
	//err := validate.Struct(param)
	//if err != nil {
	//	ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, "结构体参数错误"))
	//	return
	//}
	//param, err = product.AnyAddCategory(param)
	//if err != nil {
	//	ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, param, err.Error()))
	//	return
	//}
	//ctx.JSON(response.NewResponse(response.NoErr.Code, param, response.NoErr.Msg))
	//return
}
