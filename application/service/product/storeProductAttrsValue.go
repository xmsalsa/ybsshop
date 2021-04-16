/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-11
 * Time: 下午 14:37
 */
package product

import (
	"encoding/json"
	"gorm.io/gorm"
	_package "shop/application/libs/package"
	"shop/application/models/product"
)

type PostProductAttrValues struct {
	ProductId int
	Items     []struct {
		Value  string   `json:"value"`
		Detail []string `json:"detail"`
	} `json:"items"`
	Attrs []struct {
		Detail []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"detail"`
		Pic          string  `json:"pic"`
		Price        float64 `json:"price"`
		Cost         float64 `json:"cost"`
		OtPrice      float64 `json:"ot_price"`
		VipPrice     float64 `json:"vip_price"`
		Stock        int     `json:"stock"`
		BarCode      string  `json:"bar_code"`
		Weight       float64 `json:"weight"`
		Volume       float64 `json:"volume"`
		Brokerage    int     `json:"brokerage"`
		BrokerageTwo int     `json:"brokerage_two"`
	} `json:"attrs" validate:"required" errors:"商品规格"`
}

/* 定义结构体 */
type ProductAttrCalueService struct {
	/* 错误体 */
	isErr error
	Sql   gorm.DB
}

/**
删除商品
*/
func (ser *ProductAttrCalueService) DelectProductAttrValue(ProductId int) error {
	Sql := ser.Sql.Model(product.StoreProductAttrsValue{})
	if err := Sql.Where("product_id = ? ", ProductId).Delete(&product.StoreProductAttrsValue{}).Error; err != nil {
		return err
	}
	return nil
}

/**
创建商品属性值
*/
func (ser *ProductAttrCalueService) CreateProductAttrValue(attr PostProductAttrValues, Authority _package.Authority) {
	//删除属性表
	ser.DelectProductAttrValue(attr.ProductId)
	for _, v := range attr.Attrs {
		Sql := ser.Sql.Model(product.StoreProductAttrsValue{})
		Value := product.StoreProductAttrsValue{}
		_package.StructAssign(&Value, &v)
		_package.AddStructCommon(0, &Value, Authority)
		Value.Price = _package.F2i(v.Price)
		Value.Cost = _package.F2i(v.Cost)
		Value.OtPrice = _package.F2i(v.OtPrice)
		Value.VipPrice = _package.F2i(v.VipPrice)
		Value.Image = v.Pic
		strRet, _ := json.Marshal(v.Detail)
		// json转map
		var mRet []map[string]interface{}
		json.Unmarshal(strRet, &mRet)
		Value.Suk = _package.GetMapKey("value", mRet)
		Value.Unique = _package.GetGUID().Hex()
		Value.ProductId = attr.ProductId
		if err := Sql.Create(&Value).Error; err != nil {
			ser.isErr = err
			return
		}
	}

}

func (ser *ProductAttrCalueService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
