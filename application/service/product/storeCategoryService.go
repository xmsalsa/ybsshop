/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 18:09
 */
package product

import (
	"errors"
	"shop/application/libs/easygorm"
	"shop/application/libs/logging"
	_package "shop/application/libs/package"
	"shop/application/models/product"
)

func FindStoreCategory(Id int) (product.StoreCategory, error) {
	user := product.StoreCategory{}
	if err := easygorm.GetEasyGormDb().Where("id = ?", Id).First(&user); err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func AnyAddCategory(param product.StoreCategory) (product.StoreCategory, error) {
	Sql := easygorm.GetEasyGormDb()
	if param.Pid > 0 {
		user, error := FindStoreCategory(param.Pid)
		if error != nil {
			return param, error
		}
		if user.MerchId != param.MerchId {
			return param, errors.New("此上级分类不属于商户下")
		}
		if user.Effect == 0 {
			return param, errors.New("该数据已删除不能修改")
		}
		if user.Pid != 0 {
			return param, errors.New("此上级不属于顶级分类")
		}
		if err := Sql.Save(&param); err.Error != nil {
			return param, err.Error
		}
		return param, nil
	}
	if param.Id == 0 {
		if err := Sql.Create(&param); err.Error != nil {
			return param, err.Error
		}
		return param, nil
	}
	return param, nil
}

/**
获取分类列表
*/
func GetListCategory(where map[string]string, auth map[string]interface{}, page _package.Page) (interface{}, error) {
	var count int64
	list := make([]interface{}, 0)
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreCategory{})
	Sql = _package.IntelligenceSql(auth, Sql, true)
	for kEys, val := range where {
		if kEys == "cate_name" {
			continue
		}
		Sql.Where(kEys+" = ?", val)
	}
	if where["cate_name"] != "" {
		Sql.Where(" cate_name like ? ", "%"+where["cate_name"]+"%")
	}
	var Model []product.StoreCategory
	_package.Paging(Sql, page).Find(&Model)
	ids := make([]int, 0)
	for _, s := range Model {
		ids = append(ids, s.Id)
	}
	s, _ := findStoreCategoryChildren(ids)
	for _, k := range Model {
		f := CategoryData(k)
		if _, ok := s[k.Id]; ok {
			f["children"] = s[k.Id]
		} else {
			f["children"] = make([]interface{}, 0)
		}
		list = append(list, f)
	}
	err := Sql.Count(&count).Error
	if err != nil {
		logging.ErrorLogger.Errorf("获取分类列表错误: ", err)
		return _package.List(list, 0), err
	}
	return _package.List(list, count), nil
}

func CategoryData(s product.StoreCategory) map[string]interface{} {
	sChildren := make(map[string]interface{})
	sChildren["id"] = s.Id
	sChildren["big_pic"] = s.BigPic
	sChildren["cate_name"] = s.CateName
	sChildren["is_show"] = s.IsShow
	sChildren["pic"] = s.Pic
	sChildren["pid"] = s.Pid
	sChildren["sort"] = s.Sort
	sChildren["add_time"] = _package.UnixToStr(s.CreatedAt, "2006-01-02 15:04:05")
	return sChildren
}

/**
查询子级分类
*/
func findStoreCategoryChildren(Ids []int) (map[int][]interface{}, error) {
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreCategory{})
	children := make(map[int][]interface{}, 0)
	var Model []product.StoreCategory
	Sql.Where("pid in (?)", Ids).Find(&Model)
	for _, s := range Model {
		if _, ok := children[s.Pid]; ok {
		} else {
			children[s.Pid] = make([]interface{}, 0)
		}
		children[s.Pid] = append(children[s.Pid], CategoryData(s))
	}
	return children, nil
}

/**
查询自身
*/
func findStoreCategoryList(Ids []int) (map[int]product.StoreCategory, error) {
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreCategory{})
	children := make(map[int]product.StoreCategory)
	var Model []product.StoreCategory
	Sql.Where("id in (?)", Ids).Find(&Model)
	for _, s := range Model {
		children[s.Id] = s
	}
	return children, nil
}

/**删除商品分类 **/
func DelectCategory(id int, token map[string]interface{}, deletion map[string]interface{}) (interface{}, error) {
	Sql := easygorm.GetEasyGormDb().Model(product.StoreCategory{})
	Sql = _package.IntelligenceSql(token, Sql, true)
	deletion["effect"] = _package.EFFECTD
	if err := Sql.Where("id = ? ", id).Updates(&deletion); err.Error != nil {
		return deletion, err.Error
	}
	_, _ = DelectCategoryChildren(id, token, deletion)
	return deletion, nil
}

func DelectCategoryChildren(pid int, token map[string]interface{}, deletion map[string]interface{}) (interface{}, error) {
	Sql := easygorm.GetEasyGormDb().Model(product.StoreCategory{})
	Sql = _package.IntelligenceSql(token, Sql, true)
	deletion["effect"] = _package.EFFECTD
	if err := Sql.Where("pid = ? ", pid).Updates(&deletion); err.Error != nil {
		return deletion, err.Error
	}
	return deletion, nil
}

/* 定义结构体 */
type ProductCategoryService struct {
	/* 错误体 */
	isErr error
}

type GetGoodsCategory struct {
	Value    int    `json:"value"`
	Label    string `json:"label"`
	Disabled int    `json:"disabled"`
}

/* 获取商品分类 */
func (ser *ProductCategoryService) GetGoodsCategory(Ids []int) []GetGoodsCategory {
	data := make([]GetGoodsCategory, 0)
	pro, err := findStoreCategoryList(Ids)
	if err != nil {
		ser.isErr = err
		return data
	}
	for _, s := range pro {
		Goods := GetGoodsCategory{
			Value:    s.Id,
			Label:    "|-----" + s.CateName,
			Disabled: 0,
		}
		data = append(data, Goods)
	}
	return data
}

func (ser *ProductCategoryService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
