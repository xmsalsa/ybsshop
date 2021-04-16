/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-9
 * Time: 下午 15:09
 */
package product

import (
	"gorm.io/gorm"
	"shop/application/models/product"
	"strings"
)

type PostProductAttr struct {
	ProductId int `gorm:"not null default 0 comment('商品ID') index INT(10)"`
	Type      int `gorm:"default 0 comment('活动类型 0=商品，1=秒杀，2=砍价，3=拼团') TINYINT(1)"`
	Items     []struct {
		Value  string   `json:"value"`
		Detail []string `json:"detail"`
	} `json:"items"`
}

func attrData(data map[string]string) ([]string, []string) {
	kEys := make([]string, 0)
	iTems := make([]string, 0)

	for ks, v := range data {
		kEys = append(kEys, ks)
		iTems = append(iTems, v)
	}
	return kEys, iTems
}

/* 定义结构体 */
type ProductAttrService struct {
	/* 错误体 */
	isErr error
	Sql   gorm.DB
}

/**
删除商品
*/
func (ser *ProductAttrService) DelectProductAttr(ProductId int) error {
	Sql := ser.Sql.Model(product.StoreProductAttr{})
	if err := Sql.Where("product_id = ? ", ProductId).Delete(&product.StoreProductAttr{}).Error; err != nil {
		return err
	}
	return nil
}

/**
创建商品属性值
*/
func (ser *ProductAttrService) CreateProductAttr(attr PostProductAttr) {
	//删除属性表
	ser.DelectProductAttr(attr.ProductId)
	Sql := ser.Sql
	for _, v := range attr.Items {
		StoreProductAttr := product.StoreProductAttr{
			ProductId: attr.ProductId,
			Type:      attr.Type,
		}
		StoreProductAttr.AttrName = v.Value
		StoreProductAttr.AttrValues = strings.Join(v.Detail, ",")
		if err := Sql.Create(&StoreProductAttr).Error; err != nil {
			ser.isErr = err
			return
		}
	}

}

func (ser *ProductAttrService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
