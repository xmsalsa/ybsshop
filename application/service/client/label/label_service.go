package label

import (
	"errors"
	"fmt"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	mlabel "shop/application/models/client/label"
	"shop/application/service/component"
	"time"
)

func LabelCreate(param SLabelCreate) (mlabel.Label, error) {
	label := mlabel.Label{}

	// 标签分类判断
	// 略

	// 判断重复
	condition := map[string]interface{}{
		"name": param.Name,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(label)
	utils.Build(db, condition)
	err := db.Begin().Find(&label).Error
	if err != nil {
		return label, err
	}
	if label.Id != 0 {
		return label, errors.New(fmt.Sprintf("%s%s%s", label.TableComment(), param.Name, response.DB_RECORD_EXSIT))
	}

	utils.InitModel(&label, param.UpdatedUid, param.MerchId)
	label.Type = param.Type
	label.Source = param.Source
	label.Sort = param.Sort
	label.Name = param.Name
	label.LabelcateId = param.LabelcateId
	err = utils.GetGormDb().Create(&label).Error
	if err != nil {
		return label, err
	}

	return label, nil
}

func LabelDetail(param SLabelDetail) (mlabel.Label, error) {
	label := mlabel.Label{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(label)
	utils.Build(db, condition)
	err := db.Find(&label).Error
	if err != nil {
		return label, err
	}

	return label, nil
}

func LabelUpdate(param SLabelModify) (mlabel.Label, error) {
	label := mlabel.Label{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(label)
	utils.Build(db, condition)
	err := db.Find(&label).Error
	if err != nil {
		return label, err
	}
	if label.Id == 0 {
		return label, errors.New(response.DbRecordNoExsit.Msg)
	}

	// 判断重复
	exsit := mlabel.Label{}
	condition = make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	condition["name"] = param.Name
	db = utils.GetGormDbWithModel(exsit)
	utils.Build(db, condition)
	err = db.Where("id<>?", param.Id).Find(&exsit).Error
	if err != nil {
		return label, err
	}
	if exsit.Id != 0 {
		return label, errors.New(fmt.Sprintf("%s%s%s", exsit.TableComment(), param.Name, response.DB_RECORD_EXSIT))
	}

	// 更新操作
	data := map[string]interface{}{
		"updated_at":  time.Now().Unix(),
		"updated_uid": param.UpdatedUid,
		"name":        param.Name,
	}
	if param.LabelcateId > 0 {
		data["labelcate_id"] = param.LabelcateId
	}
	err = utils.GetGormDbWithModel(label).Where("id=?", param.Id).Updates(data).Error
	if err != nil {
		return label, err
	}
	label.LabelcateId = param.LabelcateId
	label.Name = param.Name

	return label, nil
}

func LabelDel(param SLabelDetail) (mlabel.Label, error) {
	label, err := LabelDetail(param)
	if err != nil {
		return label, err
	}

	if label.Id == 0 {
		return label, errors.New(response.DbRecordNoExsit.Msg)
	}

	// 绑定的客户记录
	api_param := SClientLabelCount{
		LabelId: param.Id,
	}
	count, err := ClientLabelCount(api_param)
	if err != nil {
		return label, err
	}
	if count > 0 {
		return label, errors.New(fmt.Sprintf("%s%s", label.TableComment(), response.DB_HAS_BINDRECORD))
	}

	// 删除操作
	data := map[string]interface{}{
		"updated_at":  time.Now().Unix(),
		"updated_uid": param.UpdatedUid,
		"effect":      0,
	}
	err = utils.GetGormDbWithModel(mlabel.Label{}).Where("id=?", param.Id).Updates(data).Error
	if err != nil {
		return label, nil
	}

	return label, nil
}

func LabelCount(param SLabelAll) (int64, error) {
	condition := make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	if param.LabelcateId > 0 {
		condition["labelcate_id"] = param.LabelcateId
	}
	var count int64 = 0
	db := utils.GetGormDbWithModel(mlabel.Label{})
	utils.Build(db, condition)
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func LabelAll(param SLabelAll, page utils.Pages) ([]mlabel.Label, error) {
	var list []mlabel.Label

	condition := make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	if param.LabelcateId > 0 {
		condition["labelcate_id"] = param.LabelcateId
	}

	db := utils.GetGormDbWithModel(mlabel.Label{})
	utils.Build(db, condition)
	err := db.Limit(int(page.Limit)).Offset(int(page.Offset)).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func LabelEditBox(param SLabelDetail) (component.ComponentBox, error) {
	var box component.ComponentBox
	var err error
	labecate_list := []mlabel.LabelCategory{}
	label := mlabel.Label{}

	// 可以开启多进程
	// 获取标签
	if param.Id > 0 {
		api_param := SLabelDetail{
			Id:      param.Id,
			MerchId: param.MerchId,
		}
		label, err = LabelDetail(api_param)
		if err != nil {
			return box, err
		}
		if label.Id == 0 {
			return box, errors.New(fmt.Sprintf("%s%s", label.TableComment(), response.DB_RECORD_NOEXSIT))
		}
	}
	// 获取标签分类
	api_param1 := SLabelcateAll{
		MerchId: param.MerchId,
	}
	labecate_list, err = LabelcateAll(api_param1, utils.PagesAll())
	if err != nil {
		return box, err
	}
	if len(labecate_list) == 0 {
		lc := mlabel.LabelCategory{}
		return box, errors.New(fmt.Sprintf("%s%s", lc.TableComment(), response.DB_RECORD_NOEXSIT))
	}

	box = component.LabelComponent(labecate_list, label)
	return box, nil
}
