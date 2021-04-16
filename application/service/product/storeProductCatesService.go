/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-12
 * Time: 上午 10:43
 */
package product

import (
	"errors"
	"fmt"
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/models/product"
)

type PostProductCates struct {
	ProductId int
	CateID    []int `json:"cate_id" validate:"required" errors:"前选择分类"`
}

type productCatesService interface {
	CreateProductCates(param PostProduct, Authority _package.Authority) (product.StoreProduct, error)
	GetProductCates(where map[string]interface{}, token map[string]interface{}, page _package.Page) interface{}
	Error() string
}

/* 定义结构体 */
type ProductCatesService struct {
	/* 错误体 */
	isErr error
}

/**
删除商品
*/
func DelectProductCates(ProductId int) error {
	Sql := easygorm.GetEasyGormDb().Model(product.StoreProductCate{})
	if err := Sql.Where("product_id = ? ", ProductId).Delete(&product.StoreProductCate{}).Error; err != nil {
		return err
	}
	return nil
}

/**
创建商品属性值
*/
func (ser *ProductCatesService) CreateProductCates(attr PostProductCates, Authority _package.Authority) PostProductCates {
	//删除属性表
	DelectProductCates(attr.ProductId)
	children, _ := findStoreCategoryList(attr.CateID)
	for _, c := range children {
		if c.Pid == 0 {
			ser.isErr = errors.New("分类不属于子级分类")
			return attr
		}
		if c.MerchId != Authority.MerchId {
			ser.isErr = errors.New("分类不属于此商户下")
			return attr
		}
	}
	for _, v := range attr.CateID {
		Sql := easygorm.GetEasyGormDb().Model(product.StoreProductCate{})
		Value := product.StoreProductCate{
			CateId:    v,
			ProductId: attr.ProductId,
			CatePid:   children[v].Pid,
		}
		_package.AddStructCommon(0, &Value, Authority)
		if err := Sql.Create(&Value).Error; err != nil {
			fmt.Println(v)
		}
	}
	return attr
}

/**	获取商品关联分类 **/
func (ser *ProductCatesService) GetProductCates(where map[string]interface{}, token map[string]interface{}, page _package.Page) []int {
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreProductCate{})
	Sql = _package.IntelligenceSql(token, Sql, true)
	Sql = _package.IntelligenceSql(where, Sql, true)
	var Model []product.StoreProductCate
	_package.Paging(Sql, page).Select("product_id").Group("product_id").Find(&Model)
	ids := make([]int, 0)
	for _, s := range Model {
		ids = append(ids, s.ProductId)
	}
	return ids
}

func (ser *ProductCatesService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
