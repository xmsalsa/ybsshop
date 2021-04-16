/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-8
 * Time: 下午 15:31
 */
package product

import (
	"errors"
	"fmt"
	"shop/application/libs/core"
	"shop/application/libs/easygorm"
	"shop/application/libs/logging"
	_package "shop/application/libs/package"
	"shop/application/models/product"
	"strings"
	"sync"
)

type productService interface {
	CreateProduct(param PostProduct, Authority _package.Authority) (product.StoreProduct, error)
	Error() string
}

/* 定义结构体 */
type ProductService struct {
	/* 错误体 */
	isErr error
}

type PostProduct struct {
	Id             int           `json:"id"`
	MerchID        int           `json:"merch_id"`
	Image          string        `json:"image" validate:"required" errors:"商品图片"`
	RecommendImage string        `json:"recommend_image"`
	SliderImage    []string      `json:"slider_image" validate:"required" errors:"商品轮播图"`
	StoreName      string        `json:"store_name" validate:"required" errors:"商品名称"`
	StoreInfo      string        `json:"store_info" errors:"商品简介"`
	Keyword        string        `json:"keyword"`
	BarCode        string        `json:"bar_code"`
	CateID         []int         `json:"cate_id" validate:"required" errors:"前选择分类"`
	Price          string        `json:"price"`
	VipPrice       string        `json:"vip_price"`
	OtPrice        string        `json:"ot_price"`
	Postage        string        `json:"postage"`
	UnitName       string        `json:"unit_name" validate:"required" errors:"单位名称"`
	Sort           int           `json:"sort"`
	Sales          int           `json:"sales"`
	Stock          int           `json:"stock"`
	IsShow         int           `json:"is_show"`
	IsHot          int           `json:"is_hot"`
	IsBenefit      int           `json:"is_benefit"`
	IsBest         int           `json:"is_best"`
	IsNew          int           `json:"is_new"`
	AddTime        int           `json:"add_time"`
	IsPostage      int           `json:"is_postage"`
	IsDel          int           `json:"is_del"`
	MerUse         int           `json:"mer_use"`
	GiveIntegral   int           `json:"give_integral"`
	Cost           string        `json:"cost"`
	IsSeckill      int           `json:"is_seckill"`
	IsBargain      interface{}   `json:"is_bargain"`
	IsGood         int           `json:"is_good"`
	IsSub          []interface{} `json:"is_sub"`
	IsVip          int           `json:"is_vip"`
	Ficti          int           `json:"ficti"`
	Browse         int           `json:"browse"`
	CodePath       string        `json:"code_path"`
	SoureLink      string        `json:"soure_link"`
	VideoLink      string        `json:"video_link"`
	TempId         int           `json:"temp_id" validate:"required" errors:"运费模板id"`
	SpecType       int           `json:"spec_type"`
	Activity       []string      `json:"activity"`
	Spu            string        `json:"spu"`
	LabelID        []int         `json:"label_id"`
	CommandWord    string        `json:"command_word"`
	Coupons        []interface{} `json:"coupons"`
	Description    string        `json:"description"`
	Items          []struct {
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
	Attr struct {
		Pic          string `json:"pic"`
		VipPrice     int    `json:"vip_price"`
		Price        int    `json:"price"`
		Cost         int    `json:"cost"`
		OtPrice      int    `json:"ot_price"`
		Stock        int    `json:"stock"`
		BarCode      string `json:"bar_code"`
		Weight       int    `json:"weight"`
		Volume       int    `json:"volume"`
		Brokerage    int    `json:"brokerage"`
		BrokerageTwo int    `json:"brokerage_two"`
	} `json:"attr"`
	CouponIds []interface{} `json:"coupon_ids"`
	Header    []struct {
		Title    string `json:"title"`
		Align    string `json:"align"`
		MinWidth int    `json:"minWidth"`
		Key      string `json:"key,omitempty"`
		ID       string `json:"__id"`
		Slot     string `json:"slot,omitempty"`
	} `json:"header"`
}

func intiData(param PostProduct) product.StoreProduct {
	StoreProduct := product.StoreProduct{}
	_package.StructAssign(&StoreProduct, &param)
	StoreProduct.Price = _package.F2si(param.Price)
	StoreProduct.Cost = _package.F2si(param.Cost)
	StoreProduct.OtPrice = _package.F2si(param.OtPrice)
	StoreProduct.Postage = _package.F2si(param.Postage)
	StoreProduct.VipPrice = _package.F2si(param.VipPrice)
	StoreProduct.SliderImage = strings.Join(param.SliderImage, ",")
	StoreProduct.CateId = strings.Replace(strings.Trim(fmt.Sprint(param.CateID), "[]"), " ", ",", -1)
	return StoreProduct
}

type InitItems struct {
	Value  string   `json:"value"`
	Detail []string `json:"detail"`
}

type InitAttrs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func initAttr(param PostProduct) PostProduct {
	if len(param.Items) == 0 {
		detail := make([]string, 0)
		InitAttr := InitItems{
			Value:  "规格",
			Detail: detail,
		}
		param.Items = append(param.Items, InitAttr)
	}
	for k, v := range param.Attrs {
		if len(v.Detail) == 0 {
			param.Attrs[k].Detail = append(param.Attrs[k].Detail, InitAttrs{
				Key:   "规格",
				Value: "默认",
			})
		}
	}
	return param
}

func (ser *ProductService) CreateProduct(param PostProduct, Authority _package.Authority) product.StoreProduct {
	//初始属性值
	param = initAttr(param)
	//初始产品值
	StoreProduct := intiData(param)
	Sql := easygorm.GetEasyGormDb()
	_package.AddStructCommon(param.Id, &StoreProduct, Authority)
	if StoreProduct.Id == 0 {
		if err := Sql.Create(&StoreProduct); err.Error != nil {
			ser.isErr = err.Error
			return StoreProduct
		}
	} else {
		if err := Sql.Updates(&StoreProduct); err.Error != nil {
			ser.isErr = err.Error
			return StoreProduct
		}
	}
	//创建更新商品分类表
	cates := new(ProductCatesService)
	cates.CreateProductCates(PostProductCates{
		CateID:    param.CateID,
		ProductId: StoreProduct.Id,
	}, Authority)
	if cates.Error() != "" {
		ser.isErr = errors.New(cates.Error())
		return StoreProduct
	}
	//创建更新商品属性表
	CreateProductAttr(PostProductAttr{
		ProductId: StoreProduct.Id,
		Type:      0,
		Items:     param.Items,
	})
	//穿建更新商品属性详情表
	CreateProductAttrService(PostProductAttrService{
		Attrs:     param.Attrs,
		Items:     param.Items,
		ProductId: StoreProduct.Id,
	}, Authority)
	//创建更新商品值表
	CreateProductAttrValue(PostProductAttrValues{
		Attrs:     param.Attrs,
		Items:     param.Items,
		ProductId: StoreProduct.Id,
	}, Authority)
	return StoreProduct
}

func (ser *ProductService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}

type GetProduct struct {
	Id             int         `json:"id"`
	MerchId        int         `json:"mer_id"`
	Image          string      `json:"image"`
	RecommendImage string      `json:"recommend_image"`
	SliderImage    []string    `json:"slider_image"`
	StoreName      string      `json:"store_name"`
	StoreInfo      string      `json:"store_info"`
	Keyword        string      `json:"keyword"`
	BarCode        string      `json:"bar_code"`
	CateId         string      `json:"cate_id"`
	Price          string      `json:"price"`
	VipPrice       string      `json:"vip_price"`
	OtPrice        string      `json:"ot_price"`
	Postage        string      `json:"postage"`
	UnitName       string      `json:"unit_name"`
	Sort           int         `json:"sort"`
	Sales          int         `json:"sales"`
	Stock          string      `json:"stock"`
	IsShow         int         `json:"is_show"`
	IsHot          int         `json:"is_hot"`
	IsBenefit      int         `json:"is_benefit"`
	IsBest         int         `json:"is_best"`
	IsNew          int         `json:"is_new"`
	AddTime        int         `json:"add_time"`
	IsPostage      int         `json:"is_postage"`
	IsDel          int         `json:"is_del"`
	MerUse         int         `json:"mer_use"`
	GiveIntegral   string      `json:"give_integral"`
	Cost           string      `json:"cost"`
	IsSeckill      int         `json:"is_seckill"`
	IsBargain      interface{} `json:"is_bargain"`
	IsGood         int         `json:"is_good"`
	IsSub          int         `json:"is_sub"`
	IsVip          int         `json:"is_vip"`
	Ficti          int         `json:"ficti"`
	Browse         int         `json:"browse"`
	CodePath       string      `json:"code_path"`
	SoureLink      string      `json:"soure_link"`
	VideoLink      string      `json:"video_link"`
	TempID         int         `json:"temp_id"`
	SpecType       int         `json:"spec_type"`
	Activity       string      `json:"activity"`
	Spu            string      `json:"spu"`
	LabelID        string      `json:"label_id"`
	CommandWord    string      `json:"command_word"`
	Collect        int         `json:"collect"`
	Likes          int         `json:"likes"`
	Visitor        int         `json:"visitor"`
	CateName       string      `json:"cate_name"`
	StockAttr      bool        `json:"stock_attr"`
}

func (ser *ProductService) GetProduct(where map[string]interface{}, token map[string]interface{}, page _package.Page) interface{} {
	var count int64
	countryCapitalMap := make([]interface{}, 0)
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreProduct{})
	Sql = _package.IntelligenceSql(token, Sql, false)
	Sql = _package.IntelligenceSql(where, Sql, false)
	var Model []product.StoreProduct
	_package.Paging(Sql, page).Find(&Model)
	for _, s := range Model {
		countryCapitalMap = append(countryCapitalMap, view(s))
	}
	err := Sql.Count(&count).Error
	if err != nil {
		logging.ErrorLogger.Errorf("获取sku模板列表错误: ", err)
		return _package.List(countryCapitalMap, 0)
	}
	return _package.List(countryCapitalMap, count)
}

func view(s product.StoreProduct) GetProduct {
	sc := GetProduct{}
	_package.StructAssign(&sc, &s)
	sc.CateName = ""
	sc.Cost = _package.IF64Str(s.Cost, 100)
	sc.GiveIntegral = _package.IF64Str(s.GiveIntegral, 100)
	sc.OtPrice = _package.IF64Str(s.OtPrice, 100)
	sc.Postage = _package.IF64Str(s.Postage, 100)
	sc.Price = _package.IF64Str(s.Price, 100)
	sc.VipPrice = _package.IF64Str(s.VipPrice, 100)
	sc.Image = _package.ImgLinkUrl(s.Image)
	sc.SliderImage = strings.Split(s.SliderImage, ",")
	return sc
}

func (ser *ProductService) ListProduct(Ids []int) []interface{} {
	Sql := easygorm.GetEasyGormDb().Model(&product.StoreProduct{})
	children := make([]interface{}, 0)
	var Model []product.StoreProduct
	Sql.Where("id in (?)", Ids).Find(&Model)
	for _, s := range Model {
		children = append(children, view(s))
	}
	return children
}

func findProduct(id int) (product.StoreProduct, error) {
	product := product.StoreProduct{}
	Sql := easygorm.GetEasyGormDb().Model(product)
	if err := Sql.Where("id = ?", id).First(&product); err.Error != nil {
		return product, err.Error
	}
	return product, nil
}

type PutSetShow struct {
	Id     int `json:"id" validate:"required" errors:"商品id"`
	IsShow int `json:"is_show" validate:"required" errors:"上下架状态"`
}

/**	上架商品 **/
func (ser *ProductService) PutSetShow(id int, isShow int, Authority _package.Authority) product.StoreProduct {
	product, err := findProduct(id)
	if err != nil {
		ser.isErr = err
		return product
	}
	if product.MerchId != Authority.MerchId {
		ser.isErr = errors.New("权限不足")
		return product
	}
	Sql := easygorm.GetEasyGormDb().Model(product)
	_package.AddStructCommon(product.Id, &product, Authority)
	product.IsShow = isShow
	if err := Sql.Updates(&product); err.Error != nil {
		ser.isErr = err.Error
	}
	return product
}

type PutUnShow struct {
	Ids    []int `json:"ids" validate:"required" errors:"下架商品id"`
	IsShow int   `json:"is_show" validate:"required" errors:"上下架状态"`
}

/** 批量上下架商品 **/
func (ser *ProductService) PutUnShow(param PutUnShow, Update map[string]interface{}, token map[string]interface{}) map[string]interface{} {
	product := product.StoreProduct{}
	Sql := easygorm.GetEasyGormDb().Model(product)
	Sql = _package.IntelligenceSql(token, Sql, true)
	Update["is_show"] = param.IsShow
	if err := Sql.Where("id in (?)", param.Ids).Updates(&Update); err.Error != nil {
		ser.isErr = err.Error
	}
	return Update
}

type GetDetails struct {
	Id int `json:"id" validate:"required" errors:"商品id"`
}

type GoodsDetailsTem struct {
	TempList    []interface{}      `json:"tempList"`
	CateList    []GetGoodsCategory `json:"cateList"`
	ProductInfo struct {
		Id             int           `json:"id"`
		MerchId        int           `json:"mer_id"`
		Image          string        `json:"image"`
		RecommendImage string        `json:"recommend_image"`
		SliderImage    []string      `json:"slider_image"`
		StoreName      string        `json:"store_name"`
		StoreInfo      string        `json:"store_info"`
		Keyword        string        `json:"keyword"`
		BarCode        string        `json:"bar_code"`
		CateId         []string      `json:"cate_id"`
		Price          string        `json:"price"`
		VipPrice       string        `json:"vip_price"`
		OtPrice        string        `json:"ot_price"`
		Postage        string        `json:"postage"`
		UnitName       string        `json:"unit_name"`
		Sort           int           `json:"sort"`
		Sales          int           `json:"sales"`
		Stock          int           `json:"stock"`
		IsShow         int           `json:"is_show"`
		IsHot          int           `json:"is_hot"`
		IsBenefit      int           `json:"is_benefit"`
		IsBest         int           `json:"is_best"`
		IsNew          int           `json:"is_new"`
		AddTime        int           `json:"add_time"`
		IsPostage      int           `json:"is_postage"`
		IsDel          int           `json:"is_del"`
		MerUse         int           `json:"mer_use"`
		GiveIntegral   int           `json:"give_integral"`
		Cost           string        `json:"cost"`
		IsSeckill      int           `json:"is_seckill"`
		IsBargain      interface{}   `json:"is_bargain"`
		IsGood         int           `json:"is_good"`
		IsSub          []interface{} `json:"is_sub"`
		IsVip          int           `json:"is_vip"`
		Ficti          int           `json:"ficti"`
		Browse         int           `json:"browse"`
		CodePath       string        `json:"code_path"`
		SoureLink      string        `json:"soure_link"`
		VideoLink      string        `json:"video_link"`
		TempId         int           `json:"temp_id"`
		SpecType       int           `json:"spec_type"`
		Activity       []interface{} `json:"activity"`
		Spu            string        `json:"spu"`
		LabelId        []string      `json:"label_id"`
		CommandWord    string        `json:"command_word"`
		Coupons        []interface{} `json:"coupons"`
		Description    string        `json:"description"`
		Items          []struct {
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
			Cost         int     `json:"cost"`
			OtPrice      int     `json:"ot_price"`
			VipPrice     int     `json:"vip_price"`
			Stock        int     `json:"stock"`
			BarCode      string  `json:"bar_code"`
			Weight       int     `json:"weight"`
			Volume       int     `json:"volume"`
			Brokerage    int     `json:"brokerage"`
			BrokerageTwo int     `json:"brokerage_two"`
		} `json:"attrs"`
		Attr struct {
			Pic          string `json:"pic"`
			VipPrice     int    `json:"vip_price"`
			Price        int    `json:"price"`
			Cost         int    `json:"cost"`
			OtPrice      int    `json:"ot_price"`
			Stock        int    `json:"stock"`
			BarCode      string `json:"bar_code"`
			Weight       int    `json:"weight"`
			Volume       int    `json:"volume"`
			Brokerage    int    `json:"brokerage"`
			BrokerageTwo int    `json:"brokerage_two"`
		} `json:"attr"`
	} `json:"productInfo"`
}

/** 获取商品详情 **/
func (ser *ProductService) GetDetails(param GetDetails) GoodsDetailsTem {
	Tem := GoodsDetailsTem{}
	product, err := findProduct(param.Id)
	if err != nil {
		ser.isErr = err
		return Tem
	}
	Dispatcher := new(core.Dispatcher)
	Dispatcher.AddEventListener("test", getGoodsAttrResult)
	Dispatcher.AddEventListener("test", getGoodsCategory)
	var wg sync.WaitGroup
	Event := core.Event{
		Wg:   wg,
		Name: "test",
	}
	Event.AddParams("id", product.Id)
	Event.AddParams("cate_id", product.CateId)
	Dispatcher.DispatchEvent(&Event)
	Event.Wg.Wait()
	if Event.Error() != "" {
		ser.isErr = errors.New(Event.Error())
		return Tem
	}
	Attrs := GetGoodsAttrResult{}
	Event.GetData("attr", &Attrs)
	Cat := make([]GetGoodsCategory, 0)
	Event.GetData("cat", &Cat)
	//Cat := Service.GetGoodsCategory(_package.SlicesStrFInt(strings.Split(product.CateId, ",")))
	Tem.CateList = Cat
	Tem.ProductInfo.Attrs = Attrs.Value
	Tem.ProductInfo.Items = Attrs.Attr
	_package.StructAssign(&Tem.ProductInfo, &product)
	Tem.ProductInfo.Price = fmt.Sprintf("%.2f", float64(product.Price/100))
	Tem.ProductInfo.Cost = fmt.Sprintf("%.2f", float64(product.Cost/100))
	Tem.ProductInfo.OtPrice = fmt.Sprintf("%.2f", float64(product.OtPrice/100))
	Tem.ProductInfo.Postage = fmt.Sprintf("%.2f", float64(product.Postage/100))
	Tem.ProductInfo.VipPrice = fmt.Sprintf("%.2f", float64(product.VipPrice/100))
	Tem.ProductInfo.SliderImage = strings.Split(product.SliderImage, ",")
	Tem.ProductInfo.CateId = strings.Split(product.CateId, ",")
	Tem.TempList = make([]interface{}, 0)
	return Tem
}

/**	获取attrs **/
func getGoodsAttrResult(enet *core.Event) {
	//defer wg.Done()
	defer enet.Wg.Done()
	if enet.Error() != "" {
		return
	}
	//获取attrs
	AttrResult := new(ProductAttrResultService)
	Attrs := AttrResult.GetGoodsAttrResult(enet.Params["id"].(int))
	if AttrResult.Error() != "" {
		enet.IsErr = errors.New(AttrResult.Error())
		return
	}
	enet.AddData("attr", Attrs)
}

/**	获取分类 **/
func getGoodsCategory(enet *core.Event) {
	defer enet.Wg.Done()
	if enet.Error() != "" {
		return
	}
	//获取分类
	Service := new(ProductCategoryService)
	Cat := Service.GetGoodsCategory(_package.SlicesStrFInt(strings.Split(enet.Params["cate_id"].(string), ",")))
	if Service.Error() != "" {
		enet.IsErr = errors.New(Service.Error())
		return
	}
	enet.AddData("cat", Cat)
}
