/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 14:13:10
 */
package label

import (
	"errors"
	"fmt"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	mlabel "shop/application/models/client/label"
)

// 批量增加多个客户多个标签
func SaveSetLabels(param SSaveSetLabels) (int, error) {
	var count int = 0
	var err error
	var list []mlabel.ClientLabelRelation

	// 整合要插入的记录
	var combination []mlabel.ClientLabelRelation
	_init := mlabel.ClientLabelRelation{}
	utils.InitModel(&_init, param.UpdatedUid, param.MerchId)
	var temp mlabel.ClientLabelRelation
	for _, client_id := range param.ClientIds {
		for _, label_id := range param.LabelIds {
			temp = _init
			temp.ClientId = client_id
			temp.LabelId = label_id
			combination = append(combination, temp)
		}
	}

	// 查找这些客户的所有记录, 然后过滤已经存在的, 剩下的再添加
	condition := map[string]interface{}{
		"client_id": param.ClientIds,
		"label_id":  param.LabelIds,
	}
	if len(param.ClientIds) == 1 {
		condition["client_id"] = param.ClientIds[0]
	}
	if len(param.LabelIds) == 1 {
		condition["label_id"] = param.LabelIds[0]
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(mlabel.ClientLabelRelation{})
	utils.Build(db, condition)
	err = db.Order("client_id ASC").Find(&list).Error
	if err != nil {
		return count, err
	}
	// 去重
	for _, l := range list {
		for k, v := range combination {
			if v.ClientId == l.ClientId && v.LabelId == l.LabelId {
				combination = append(combination[:k], combination[k+1:]...)
				continue
			}
		}
	}
	// 插入记录
	for _, l := range combination {
		err = db.Create(&l).Error
		if err != nil {
			return count, err
		}
		count += 1
	}

	return count, nil
}

func SetClientLabel(param SSetClientLabel) (map[string]int, error) {
	var del_count, add_count int = 0, 0
	var err error
	data := map[string]int{
		"del": 0,
		"add": 0,
	}

	// 获取客户的标签列表
	var relation_list []mlabel.ClientLabelRelation
	var del_relation_ids, binded_label_ids []int // 解绑要删除的关联记录ID, 已经绑定的标签ID
	condition := map[string]interface{}{
		"client_id": param.ClientId,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(mlabel.ClientLabelRelation{})
	utils.Build(db, condition)
	err = db.Find(&relation_list).Error
	if err != nil {
		return data, err
	}
	for _, v := range relation_list {
		if utils.InArrayInt(param.LabelIds, v.LabelId) {
			binded_label_ids = append(binded_label_ids, v.LabelId)
			continue
		}
		del_relation_ids = append(del_relation_ids, v.Id)
	}

	// 删除
	del_count = len(del_relation_ids)
	if del_count > 0 {
		delData := utils.InitUpdate(param.UpdatedUid)
		delData["effect"] = 0
		err = utils.GetGormDbWithModel(mlabel.ClientLabelRelation{}).Where(del_relation_ids).Updates(delData).Error
		if err != nil {
			return data, nil
		}
	}
	// 增加
	if len(param.LabelIds) > len(binded_label_ids) {
		_init := mlabel.ClientLabelRelation{}
		utils.InitModel(&_init, param.UpdatedUid, param.MerchId)
		db = utils.GetGormDbWithModel(mlabel.ClientLabelRelation{})
		for _, v := range param.LabelIds {
			if utils.InArrayInt(binded_label_ids, v) {
				continue
			}
			relation := mlabel.ClientLabelRelation{}
			relation = _init
			relation.ClientId = param.ClientId
			relation.LabelId = v
			err = db.Create(&relation).Error
			if err != nil {
				return data, nil
			}
			add_count += 1
		}
	}
	data["add"] = add_count
	data["del"] = del_count
	return data, nil
}

// 获取标签绑定的客户数量
func ClientLabelCount(param SClientLabelCount) (int64, error) {
	var clientlabel = mlabel.ClientLabelRelation{}
	var count int64 = 0

	condition := make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	if param.ClientId > 0 {
		condition["client_id"] = param.ClientId
	}
	if param.LabelId > 0 {
		condition["label_id"] = param.LabelId
	}
	if len(condition) == 0 {
		return count, errors.New(fmt.Sprintf("%s%s", clientlabel.TableComment(), response.VALIDATOR_FAIL))
	}

	db := utils.GetGormDbWithModel(clientlabel)
	utils.Build(db, condition)
	err := db.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

// 客户-标签分类-标签 树型
func ClientLabelTree(param SClientLabelTree) ([]map[string]interface{}, error) {
	var tree = make([]map[string]interface{}, 0)

	// 获取客户绑定的标签
	var list []mlabel.ClientLabelRelation
	var label_ids []int // 客户绑定的标签ID
	condition := map[string]interface{}{
		"client_id": param.ClientId,
	}
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	db := utils.GetGormDbWithModel(mlabel.ClientLabelRelation{})
	utils.Build(db, condition)
	err := db.Find(&list).Error
	if err != nil {
		return tree, err
	}
	if len(list) == 0 {
		return tree, nil
	}
	for _, v := range list {
		label_ids = append(label_ids, v.LabelId)
	}

	// 通过标签ID 获取标签分类ID
	delete(condition, "client_id")
	var clientlabel_list []mlabel.Label
	var labelcate_ids []int
	condition["id"] = label_ids
	db = utils.GetGormDbWithModel(mlabel.Label{})
	utils.Build(db, condition)
	err = db.Group("labelcate_id").Find(&clientlabel_list).Error
	if err != nil {
		return tree, err
	}
	for _, v := range clientlabel_list {
		labelcate_ids = append(labelcate_ids, v.LabelcateId)
	}

	// 用于获取标签分类及标签列表
	page := utils.PagesAll()

	// 获取标签分类列表
	api_param := SLabelcateAll{
		Ids: labelcate_ids,
	}
	if param.MerchId > 0 {
		api_param.MerchId = param.MerchId
	}
	labelcate_list, err1 := LabelcateAll(api_param, page)
	if err1 != nil {
		return tree, err1
	}

	// 获取标签列表
	api_param1 := SLabelAll{
		LabelcateIds: labelcate_ids,
	}
	if param.MerchId > 0 {
		api_param1.MerchId = param.MerchId
	}
	label_list, err2 := LabelAll(api_param1, page)
	if err2 != nil {
		return tree, err2
	}

	// var cate_temp, label_temp map[string]interface{}
	for _, cate := range labelcate_list {
		cate_temp := cate.RepsToDesc()                // 分类
		children := make([]map[string]interface{}, 0) // 子级, 即 label
		for k, v := range label_list {
			label_temp := v.RepsToDesc()
			label_temp["disabled"] = false
			if utils.InArrayInt(label_ids, v.Id) {
				label_temp["disabled"] = true
			}
			if v.LabelcateId == cate.Id {
				children = append(children, label_temp)
				label_list = append(label_list[:k], label_list[k+1:]...)
			}
		}
		fmt.Println("--------label_list-------")
		fmt.Println(label_list)
		cate_temp["label"] = children
		tree = append(tree, cate_temp)
	}

	return tree, nil
}
