package label

import (
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	slabel "shop/application/service/client/label"

	"github.com/kataras/iris/v12"
)

func LabelcateCreate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelcateCreate{
		Id:         0,
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

	labelcate, err1 := slabel.LabelcateCreate(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": labelcate.Id}))
	return
}

func LabelcateDetail(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	// 参数处理
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		ctx.JSON(response.RespValidatorFail("id未提交或不是整数"))
		return
	}
	// 参数处理
	param := slabel.SLabelcateDetail{
		Id:      id,
		MerchId: authority.MerchId,
	}

	labelcate, err1 := slabel.LabelcateDetail(param)
	if err1 != nil {
		ctx.JSON(response.RespDbExecutionFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(labelcate.RepsToDesc()))
	return
}

func LabelcatePages(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	page, err := utils.PageInit(ctx)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(err.Error()))
		return
	}

	// 组合参数
	param := slabel.SLabelcateAll{
		MerchId: authority.MerchId,
	}
	var data = make(map[string]interface{})
	var count int64
	var list = make([]interface{}, 0)

	// 获取数量
	count, err = slabel.LabelcateCount(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	// 列表
	if count > 0 {
		labelcateList, err := slabel.LabelcateAll(param, page)
		if err != nil {
			ctx.JSON(response.RespFail(err.Error()))
			return
		}
		for _, v := range labelcateList {
			list = append(list, v.RepsToDesc())
		}
	}

	data["count"] = count
	data["list"] = list
	ctx.JSON(response.RespSuccess(data))
	return
}

func LabelcateUpdate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelcateUpdate{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	_, err = slabel.LabelcateUpdate(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": int(param.Id)}))
	return
}

func LabelcateDel(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := slabel.SLabelcateDetail{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	_, err = slabel.LabelcateDel(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": param.Id}))
	return
}

func LabelcateEditBox(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	var err error
	// 参数处理
	param := slabel.SLabelcateDetail{
		Id:         ctx.URLParamIntDefault("id", 0),
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err = utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	box, err := slabel.LabelcateEditBox(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(box))
	return
}
