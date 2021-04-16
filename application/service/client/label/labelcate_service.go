/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-07 21:59:27
 */
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

/*
 标签分类创建
*/
func LabelcateCreate(param SLabelcateCreate) (mlabel.LabelCategory, error) {
	labelcate := mlabel.LabelCategory{}

	// 判断重名
	condition := map[string]interface{}{
		"merch_id": param.MerchId,
		"name":     param.Name,
	}
	exsit := mlabel.LabelCategory{}
	db := utils.GetGormDbWithModel(mlabel.LabelCategory{})
	utils.Build(db, condition)
	err := db.Find(&exsit).Error
	if err != nil {
		return exsit, err
	}
	if exsit.Id != 0 {
		return labelcate, errors.New(fmt.Sprintf("%s%s%s", exsit.TableComment(), param.Name, response.DB_RECORD_EXSIT))
	}

	utils.InitModel(&labelcate, param.UpdatedUid, param.MerchId)
	labelcate.Name = param.Name
	labelcate.Sort = param.Sort
	err = utils.GetGormDbWithModel(mlabel.LabelCategory{}).Create(&labelcate).Error
	if err != nil {
		return labelcate, err
	}

	return labelcate, nil
}

/*
 标签分类详情
*/
func LabelcateDetail(param SLabelcateDetail) (mlabel.LabelCategory, error) {
	labelcate := mlabel.LabelCategory{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(labelcate)
	utils.Build(db, condition)
	err := db.Find(&labelcate).Error
	if err != nil {
		return labelcate, err
	}

	return labelcate, nil
}

/*
 标签分类列表
 @user unknown
 @date 2021-04-07
*/
func LabelcateAll(param SLabelcateAll, page utils.Pages) ([]mlabel.LabelCategory, error) {
	var list []mlabel.LabelCategory

	condition := make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	if len(param.Ids) > 0 {
		condition["id"] = param.Ids
	}

	db := utils.GetGormDbWithModel(mlabel.LabelCategory{})
	utils.Build(db, condition)
	err := db.Order("sort DESC").Limit(int(page.Limit)).Offset(int(page.Offset)).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

/*
 标签分类数量
*/
func LabelcateCount(param SLabelcateAll) (int64, error) {
	var count int64

	condition := make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(mlabel.LabelCategory{})
	utils.Build(db, condition)
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

/*
 标签分类更新
 @param obj 条件
 @param data 更新数据
*/
func LabelcateUpdate(param SLabelcateUpdate) (mlabel.LabelCategory, error) {
	labelcate := mlabel.LabelCategory{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(labelcate)
	utils.Build(db, condition)
	err := db.Find(&labelcate).Error
	if err != nil {
		return labelcate, err
	}
	if labelcate.Id == 0 {
		return labelcate, errors.New(response.DbRecordNoExsit.Msg)
	}

	// 判断重复
	exsit := mlabel.LabelCategory{}
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
		"sort":        param.Sort,
	}
	err = utils.GetGormDbWithModel(mlabel.LabelCategory{}).Where("id=?", param.Id).Updates(data).Error
	if err != nil {
		return labelcate, err
	}
	labelcate.Sort = param.Sort
	labelcate.Name = param.Name

	return labelcate, nil
}

/*
 标签分类删除
*/
func LabelcateDel(param SLabelcateDetail) (mlabel.LabelCategory, error) {
	labelcate := mlabel.LabelCategory{}

	condition := map[string]interface{}{
		"id": param.Id,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(labelcate)
	utils.Build(db, condition)
	err := db.Find(&labelcate).Error
	if err != nil {
		return labelcate, err
	}
	if labelcate.Id == 0 {
		return labelcate, errors.New(response.DbRecordNoExsit.Msg)
	}

	// 分类下标签数量
	api_param := SLabelAll{
		LabelcateId: param.Id,
		MerchId:     param.MerchId,
	}
	var count int64 = 0
	count, err = LabelCount(api_param)
	if err != nil {
		return labelcate, err
	}
	if count > 0 {
		return labelcate, errors.New(fmt.Sprintf("%s %s%s", labelcate.TableComment(), new(mlabel.Label).TableComment(), response.DB_HAS_RECORD))
	}

	// 删除操作
	data := utils.InitUpdate(param.UpdatedUid)
	data["effect"] = 0
	fmt.Printf("%T, %v", data, data)
	err = utils.GetGormDbWithModel(mlabel.LabelCategory{}).Where("id=?", param.Id).Updates(data).Error
	if err != nil {
		return labelcate, err
	}
	labelcate.Effect = 0

	return labelcate, nil
}

func LabelcateEditBox(param SLabelcateDetail) (component.ComponentBox, error) {
	var box component.ComponentBox
	var err error
	labelcate := mlabel.LabelCategory{}

	if param.Id > 0 {
		api_param := SLabelcateDetail{
			Id:      param.Id,
			MerchId: param.MerchId,
		}
		labelcate, err = LabelcateDetail(api_param)
		if err != nil {
			return box, err
		}
		if labelcate.Id == 0 {
			return box, errors.New(fmt.Sprintf("%s%s", labelcate.TableComment(), response.DB_RECORD_NOEXSIT))
		}
	}

	box = component.LabelcateComponent(labelcate)
	return box, nil
}
