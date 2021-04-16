package cgroup

import (
	_package "shop/application/libs/package"
	"shop/application/libs/response"
	"shop/application/libs/utils"

	sgroup "shop/application/service/client/group"

	"github.com/kataras/iris/v12"
)

func GroupCreate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	// 参数处理
	param := sgroup.SGroupCreate{
		Id:         0,
		Sort:       1,
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
	group, err1 := sgroup.GroupCreate(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": group.Id}))
	return
}
func GroupDetail(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		ctx.JSON(response.RespValidatorFail("id未提交或不是整数"))
		return
	}
	param := sgroup.SGroupDetail{
		Id:      id,
		MerchId: authority.MerchId,
	}

	group, err := sgroup.GroupDetail(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(group.RepsToDesc()))
	return
}
func GroupPages(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	page, err := utils.PageInit(ctx)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(err.Error()))
		return
	}

	// 组合参数
	param := sgroup.SGroupAll{
		MerchId: authority.MerchId,
	}

	var data = make(map[string]interface{})
	var count int64
	var list = make([]interface{}, 0)

	// 获取数量
	count, err = sgroup.GroupCount(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	// 列表
	if count > 0 {
		grouplist, err1 := sgroup.GroupPages(param, page)
		if err1 != nil {
			ctx.JSON(response.RespFail(err1.Error()))
			return
		}
		for _, item := range grouplist {
			list = append(list, item.RepsToDesc())
		}
	}

	data["count"] = count
	data["list"] = list
	ctx.JSON(response.RespSuccess(data))
	return
}
func GroupUpdate(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := sgroup.SGroupUpdate{
		Sort:       1,
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	_, err = sgroup.GroupUpdate(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": int(param.Id)}))
	return
}
func GroupDel(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	// 参数处理
	param := sgroup.SGroupDetail{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}
	ctx.ReadJSON(&param)
	err := utils.Validate.Struct(param)
	if err != nil {
		ctx.JSON(response.RespValidatorFail(utils.ValidErrMsg(err)))
		return
	}

	_, err = sgroup.GroupDel(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(map[string]int{"id": param.Id}))
	return
}
func SetGroup(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)

	param := sgroup.SSetGroup{
		MerchId:    authority.MerchId,
		UpdatedUid: authority.UserId,
	}

	err := sgroup.SetGroup(param)
	if err != nil {
		ctx.JSON(response.RespFail(err.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(param))
	return
}
func GroupEditBox(ctx iris.Context) {
	authority := _package.GetAuthority(ctx)
	// 参数处理
	param := sgroup.SGroupDetail{
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

	box, err1 := sgroup.GroupEditBox(param)
	if err1 != nil {
		ctx.JSON(response.RespFail(err1.Error()))
		return
	}

	ctx.JSON(response.RespSuccess(box))
	return
}
