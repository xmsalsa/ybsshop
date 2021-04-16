package group

import (
	"errors"
	"fmt"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	mgroup "shop/application/models/client/group"
	"shop/application/service/component"
	"time"
)

func GroupCreate(param SGroupCreate) (mgroup.ClientGroup, error) {
	group := mgroup.ClientGroup{}

	// 判断重复
	condition := map[string]interface{}{
		"name": param.Name,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(group)
	utils.Build(db, condition)
	err := db.Find(&group).Error
	if err != nil {
		return group, err
	}
	if group.Id != 0 {
		return group, errors.New(fmt.Sprintf("%s%s%s", group.TableComment(), param.Name, response.DB_RECORD_EXSIT))
	}

	utils.InitModel(&group, param.UpdatedUid, param.MerchId)
	group.Name = param.Name
	group.Sort = 1
	err = utils.GetGormDbWithModel(mgroup.ClientGroup{}).Create(&group).Error
	if err != nil {
		return group, err
	}

	return group, nil
}
func GroupDetail(param SGroupDetail) (mgroup.ClientGroup, error) {
	group := mgroup.ClientGroup{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(group)
	utils.Build(db, condition)
	err := db.Find(&group).Error
	if err != nil {
		return group, err
	}

	return group, nil
}
func GroupCount(param SGroupAll) (int64, error) {
	var count int64

	condition := map[string]interface{}{
		"merch_id": param.MerchId,
	}

	db := utils.GetGormDbWithModel(mgroup.ClientGroup{})
	utils.Build(db, condition)
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
func GroupPages(param SGroupAll, page utils.Pages) ([]mgroup.ClientGroup, error) {
	var list []mgroup.ClientGroup

	condition := map[string]interface{}{
		"merch_id": param.MerchId,
	}
	db := utils.GetGormDbWithModel(mgroup.ClientGroup{})
	utils.Build(db, condition)
	err := db.Limit(int(page.Limit)).Offset(int(page.Offset)).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
func GroupUpdate(param SGroupUpdate) (mgroup.ClientGroup, error) {
	group := mgroup.ClientGroup{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(group)
	utils.Build(db, condition)
	err := db.Find(&group).Error
	if err != nil {
		return group, err
	}
	if group.Id == 0 {
		return group, errors.New(response.DbRecordNoExsit.Msg)
	}

	// 判断重复
	exsit := mgroup.ClientGroup{}
	condition = make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	condition["name"] = param.Name
	db = utils.GetGormDbWithModel(exsit)
	utils.Build(db, condition)
	err = db.Where("id<>?", param.Id).Find(&exsit).Error
	if err != nil {
		return exsit, err
	}
	if exsit.Id != 0 {
		return exsit, errors.New(fmt.Sprintf("%s%s%s", exsit.TableComment(), param.Name, response.DB_RECORD_EXSIT))
	}

	// 更新操作
	data := map[string]interface{}{
		"updated_at":  time.Now().Unix(),
		"updated_uid": param.UpdatedUid,
		"name":        param.Name,
		//"sort":        param.Sort,
	}
	err = utils.GetGormDbWithModel(mgroup.ClientGroup{}).Where("id=?", param.Id).Updates(data).Error
	if err != nil {
		return group, err
	}
	group.Name = param.Name

	return group, nil
}
func GroupDel(param SGroupDetail) (mgroup.ClientGroup, error) {
	group := mgroup.ClientGroup{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(group)
	utils.Build(db, condition)
	err := db.Find(&group).Error
	if err != nil {
		return group, err
	}
	if group.Id == 0 {
		return group, errors.New(response.DbRecordNoExsit.Msg)
	}

	// 绑定分组的客户
	// 判断分组是否有绑定客户, 有绑定不能删除 后续补上

	// 删除操作
	data := utils.InitUpdate(param.UpdatedUid)
	data["effect"] = 0
	fmt.Printf("%T, %v", data, data)
	err = utils.GetGormDbWithModel(mgroup.ClientGroup{}).Where("id=?", param.Id).Updates(data).Error
	if err != nil {
		return group, err
	}
	group.Effect = 0

	return group, nil
}
func SetGroup(param SSetGroup) error {

	return nil
}

//增加/编辑分组时, 电脑端的组件
func GroupEditBox(param SGroupDetail) (component.ComponentBox, error) {
	var box component.ComponentBox
	var err error
	group := mgroup.ClientGroup{}

	if param.Id > 0 {
		api_param := SGroupDetail{
			Id:      param.Id,
			MerchId: param.MerchId,
		}
		group, err = GroupDetail(api_param)
		if err != nil {
			return box, err
		}
		if group.Id == 0 {
			return box, errors.New(fmt.Sprintf("%s%s", group.TableComment(), response.DB_RECORD_NOEXSIT))
		}
	}

	box = component.GroupComponent(group)
	return box, nil
}
