/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-16
 * Time: 上午 9:20
 */
package product

import (
	"gorm.io/gorm"
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/models/product"
)

/* 定义结构体 */
type ProductDescriptionsService struct {
	/* 错误体 */
	Sql   gorm.DB
	isErr error
}

/**
删除商品
*/
func (ser *ProductDescriptionsService) delectProductDescription(ProductId int) error {
	Sql := ser.Sql.Model(product.StoreProductDescription{})
	if err := Sql.Where("product_id = ? ", ProductId).Delete(&product.StoreProductDescription{}).Error; err != nil {
		return err
	}
	return nil
}

type CreateProductDescriptionService struct {
	ProductId   int    `gorm:"not null default 0 comment('商品ID') index INT(10)"`
	Description string `json:"description"`
	Type        int    `gorm:"default 0 comment('活动类型 0=商品，1=秒杀，2=砍价，3=拼团') TINYINT(1)"`
	Authority   _package.Authority
}

/**	创建商品详情 **/
func (ser *ProductDescriptionsService) CreateProductDescriptionService(attr CreateProductDescriptionService) {
	//删除属性表
	err := ser.delectProductDescription(attr.ProductId)
	if err != nil {
		ser.isErr = err
		return
	}
	Sql := ser.Sql.Model(product.StoreProductDescription{})
	StoreProductAttr := product.StoreProductDescription{
		ProductId:   attr.ProductId,
		Description: attr.Description,
		Type:        attr.Type,
	}
	if err := Sql.Create(&StoreProductAttr).Error; err != nil {
		ser.isErr = err
		return
	}
}

/**	获取商品详情 */
func (ser *ProductDescriptionsService) GetProductDescriptionService(ProductId int) string {
	attr := product.StoreProductDescription{}
	if err := easygorm.GetEasyGormDb().Where("product_id = ?", ProductId).First(&attr); err.Error != nil {
		ser.isErr = err.Error
		return ""
	}
	return attr.Description
}

func (ser *ProductDescriptionsService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
