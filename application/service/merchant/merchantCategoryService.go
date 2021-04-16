/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 14:42
 */
package merchant

import (
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/models/merchant"
)

type merchantCategoryService interface {
}

/* 定义结构体 */
type MerchantCategoryService struct {
	/* 错误体 */
	isErr error
}
type GetMerchantCategorylst struct {
	Id             int    `json:"id"`
	CommissionRate string `json:"commission_rate"`
	CategoryName   string `json:"category_name"`
	Config         string `gorm:"not null comment('商户分类名称')"`
}

func (ser *MerchantCategoryService) GetMerchantCategorylst(token map[string]interface{}, page _package.Page) interface{} {
	var count int64
	Sql := easygorm.GetEasyGormDb().Model(&merchant.MerchantsCategory{})
	//Sql = _package.IntelligenceSql(token, Sql, true)
	var Model []GetMerchantCategorylst
	_package.Paging(Sql, page).Find(&Model)
	for k, s := range Model {
		Model[k].CommissionRate = s.CommissionRate + "%"
	}
	err := Sql.Count(&count).Error
	if err != nil {
		return _package.List(Model, 0)
	}
	return _package.List(Model, count)
}

func (ser *MerchantCategoryService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
