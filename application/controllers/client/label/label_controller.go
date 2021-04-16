package label

import (
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	slabel "shop/application/service/client/label"

	"github.com/kataras/iris/v12"
)

func LabelCreate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelCreate{
		Id:         0,
		Source:     1,
		Sort:       1,
		Type:       1,
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	// 创建
	label, err1 := slabel.LabelCreate(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": label.Id}))
	return
}

func LabelDetail(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		ctx.JSON(response.RespValidatorFail("id未提交或不是整数"))
		return
	}
	param := slabel.SLabelDetail{
		Id:      id,
		MerchId: authority.MerchId,
	}

	label, err := slabel.LabelDetail(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(label.RepsToDesc()))
	return
}

func LabelPages(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	// 参数处理
	labelcate_id := ctx.URLParamIntDefault("labelcate_id", 0)
	// if labelcate_id == 0 {
	// 	ctx.JSON(response.RespValidatorFail("labelcate_id 未提交或不是整数"))
	// 	return
	// }

	page, err := utils.PageInit(ctx)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(err.Error()))
		return
	}

	param := slabel.SLabelAll{
		MerchId: authority.MerchId,
	}
	if labelcate_id > 0 {
		param.LabelcateId = labelcate_id
	}

	var data = make(map[string]interface{})
	var count int64
	var list = make([]interface{}, 0)

	// 获取数量
	count, err = slabel.LabelCount(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	// 列表
	if count > 0 {
		all, err := slabel.LabelAll(param, page)
		if err != nil {
			ctx.JSON(response.RespFail(err.Error()))
			return
		}

		for _, l := range all {
			list = append(list, l.RepsToDesc())
		}
	}

	data["count"] = count
	data["list"] = list
	ctx.JSON(response.RespSuccess(data))
	return
}

func LabelUpdate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelModify{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	_, err1 := slabel.LabelUpdate(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": param.Id}))
	return
}

func LabelDel(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelDetail{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	_, err1 := slabel.LabelDel(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": param.Id}))
	return
}

func LabelEditBox(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelDetail{
		Id:         ctx.URLParamIntDefault("id", 0),
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	box, err1 := slabel.LabelEditBox(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(box))
	return
}
