/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 21:20:18
 */
package label

import (
	"fmt"
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	"shop/application/service/client/label"

	"github.com/kataras/iris/v12"
)

// 批量增加多个客户多个标签
func SaveSetLabels(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := label.SSaveSetLabels{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	var count int = 0
	count, err = label.SaveSetLabels(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"count": count}))
	return
}

// 编辑单客户标签(增/删)
func SetClientLabel(ctx iris.Context) {
	var data map[string]int
	var err error
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := label.SSetClientLabel{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err = utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}
	fmt.Println(param)

	data, err = label.SetClientLabel(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(data))
	return
}

// 客户-标签分类-标签 树型
func ClientLabelTree(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	client_id := ctx.URLParamIntDefault("client_id", 0)
	param := label.SClientLabelTree{
		MerchId:  authority.MerchId,
		ClientId: client_id,
	}
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	tree, err := label.ClientLabelTree(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]interface{}{"tree": tree}))
	return
}
