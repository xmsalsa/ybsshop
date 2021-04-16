/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-9
 * Time: 下午 17:46
 */
package product

import (
	"encoding/json"
	"fmt"
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/models/product"
	"time"
)

type PostProductAttrService struct {
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

func findStoreProductAttr(productId int) (product.StoreProductAttrsResult, error) {
	attr := product.StoreProductAttrsResult{}
	if err := easygorm.GetEasyGormDb().Where("product_id = ?", productId).First(&attr); err.Error != nil {
		return attr, err.Error
	}
	return attr, nil
}

/**	创建商品属性值 */
func CreateProductAttrService(attr PostProductAttrService, Authority _package.Authority) (product.StoreProductAttrsResult, error) {
	Sql := easygorm.GetEasyGormDb()
	att, s := findStoreProductAttr(attr.ProductId)
	data := make(map[string]interface{})
	data["attr"] = attr.Items
	data["value"] = attr.Attrs
	Result, _ := _package.MapToJson(data)
	if s != nil {
		att := product.StoreProductAttrsResult{
			ProductId:  attr.ProductId,
			ChangeTime: time.Now().Unix(),
			Result:     Result,
		}
		_package.AddStructCommon(0, &att, Authority)
		if err := Sql.Create(&att); err.Error != nil {
			return att, err.Error
		}
		return att, nil
	}
	att.Result = Result
	att.ChangeTime = time.Now().Unix()
	if err := Sql.Save(&att); err.Error != nil {
		return att, err.Error
	}
	return att, nil
}

/* 定义结构体 */
type ProductAttrResultService struct {
	/* 错误体 */
	isErr error
}

type GetGoodsAttrResult struct {
	Attr []struct {
		Value  string   `json:"value"`
		Detail []string `json:"detail"`
	} `json:"attr"`
	Value []struct {
		Detail []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"detail"`
		Pic          string  `json:"pic"`
		Price        float64 `json:"price"`
		Cost         int     `json:"cost"`
		OtPrice      int     `json:"ot_price"`
		VipPrice     int     `json:"vip_price"`
		Stock        int     `json:"stock"`
		BarCode      string  `json:"bar_code"`
		Weight       int     `json:"weight"`
		Volume       int     `json:"volume"`
		Brokerage    int     `json:"brokerage"`
		BrokerageTwo int     `json:"brokerage_two"`
	} `json:"value"`
}

func (ser *ProductAttrResultService) GetGoodsAttrResult(productId int) GetGoodsAttrResult {
	Service := GetGoodsAttrResult{}
	attr, err := findStoreProductAttr(productId)
	if err != nil {
		ser.isErr = err
		return Service
	}
	errs := json.Unmarshal([]byte(attr.Result), &Service)
	fmt.Println(attr.Result, Service, Service.Value)
	if errs != nil {
		ser.isErr = errs
		return Service
	}
	return Service
}

func (ser *ProductAttrResultService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
