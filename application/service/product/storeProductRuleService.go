/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 15:19
 */
package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"shop/application/libs/easygorm"
	"shop/application/libs/logging"
	_package "shop/application/libs/package"
	"shop/application/models/product"
	"strconv"
	"strings"
)

func findStoreProductRuleService(Id int) (product.StoreProductRule, error) {
	user := product.StoreProductRule{}
	if err := easygorm.GetEasyGormDb().Where("id = ?", Id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

type RuleServiceParam struct {
	Id       int    `json:"id" `
	RuleName string `json:"rule_name" validate:"required"`
	Spec     []struct {
		Value  string   `json:"value" validate:"required"`
		Detail []string `json:"detail" validate:"required"`
	} `json:"spec" validate:"required"`
}

func AnyAddRule(param product.StoreProductRule) (product.StoreProductRule, error) {
	Sql := easygorm.GetEasyGormDb()
	if param.Id > 0 {
		StoreProductRule := product.StoreProductRule{}
		user, err := findStoreProductRuleService(param.Id)
		if err != nil {
			return StoreProductRule, err
		}
		fmt.Println(user, StoreProductRule)
		if user == (StoreProductRule) {
			return StoreProductRule, errors.New("数据不存在")
		}
		if user.MerchId != param.MerchId {
			return StoreProductRule, errors.New("权限不足")
		}
		if user.Effect == 0 {
			return StoreProductRule, errors.New("该数据已删除不能修改")
		}
		if err := Sql.Save(&param).Error; err != nil {
			return param, err
		}
		return param, nil
	}
	if param.Id == 0 {
		if err := Sql.Create(&param).Error; err != nil {
			return param, err
		}
		return param, nil
	}
	return param, nil
}

func GetListRule(ruleName string, token map[string]interface{}, page _package.Page) (interface{}, string) {
	var count int64
	countryCapitalMap := make([]map[interface{}]interface{}, 0)
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreProductRule{})
	Sql = _package.IntelligenceSql(token, Sql, true)
	if ruleName != "" {
		Sql.Where(" rule_name like ? ", "%"+ruleName+"%")
	}
	var Model []product.StoreProductRule
	_package.Paging(Sql, page).Find(&Model)
	for _, s := range Model {
		mSlice := make([]struct {
			Detail []string `json:"detail"`
			Value  string   `json:"value"`
		}, 0)
		err := json.Unmarshal([]byte(s.RuleValue), &mSlice)
		if err != nil {
			return countryCapitalMap, err.Error()
		}
		attrName := ""
		attrValue := make([]string, 0)
		for sKey := range mSlice {
			if attrName != "" {
				attrName = attrName + ","
			}
			attrName = fmt.Sprintf("%s%s", attrName, mSlice[sKey].Value)
			attrValue = append(attrValue, strings.Join(mSlice[sKey].Detail, ","))
		}
		sc := make(map[interface{}]interface{})
		sc["id"] = s.Id
		sc["merch_id"] = s.MerchId
		sc["attr_name"] = attrName
		sc["attr_value"] = attrValue
		sc["rule_name"] = s.RuleName
		sc["rule_value"] = s.RuleValue
		sc["created_at"] = _package.UnixToStr(s.CreatedAt, "2006-01-02 15:04:05")
		countryCapitalMap = append(countryCapitalMap, sc)
	}
	err := Sql.Count(&count).Error
	if err != nil {
		logging.ErrorLogger.Errorf("获取sku模板列表错误: ", err)
		return _package.List(countryCapitalMap, 0), err.Error()
	}
	return _package.List(countryCapitalMap, count), ""
}

type AttrRuleDelete struct {
	All int    `json:"all"`
	Ids string `json:"ids"`
}

func DelectRule(param AttrRuleDelete, token map[string]interface{}) (AttrRuleDelete, error) {
	Sql := easygorm.GetEasyGormDb().Model(product.StoreProductRule{})
	Sql = _package.IntelligenceSql(token, Sql, true)
	LogicalDeletion := make(map[string]interface{})
	LogicalDeletion["effect"] = _package.EFFECTD
	idsString := strings.Split(param.Ids, ",")
	ids := make([]int, 0)
	for i := 0; i < len(idsString); i++ {
		intValue, err := strconv.Atoi(idsString[i])
		if err != nil {
			return param, err
		}
		ids = append(ids, intValue)
	}
	if len(ids) == 0 {
		return param, errors.New("请删除删除id")
	}
	if err := Sql.Where(ids).Updates(&LogicalDeletion); err.Error != nil {
		return param, err.Error
	}
	return param, nil
}
