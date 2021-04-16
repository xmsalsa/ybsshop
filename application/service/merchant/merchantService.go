/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 15:19
 */
package merchant

import (
	"errors"
	"shop/application/libs/easygorm"
	_package "shop/application/libs/package"
	"shop/application/models/merchant"
	"shop/application/service/system"
	"time"
)

type merchantService interface {
}

/* 定义结构体 */
type MerchantService struct {
	/* 错误体 */
	isErr error
}

type GetMerchantlst struct {
	Sort       int    `json:"sort"`
	Id         int    `json:"id"`
	MerName    string `json:"mer_name"`
	RealName   string `json:"real_name"`
	MerPhone   string `json:"mer_phone"`
	MerAddress string `json:"mer_address"`
	Mark       string `json:"mark"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	IsBest     int    `json:"is_best"`
	IsTrader   int    `json:"is_trader"`
	Admin      struct {
		MerID   int    `json:"mer_id"`
		Account string `json:"account"`
	} `json:"admin"`
}

func (ser *MerchantService) GetMerchantlst(where map[string]interface{}, token map[string]interface{}, page _package.Page) interface{} {
	var count int64
	Sql := easygorm.GetEasyGormDb().Model(&merchant.Merchant{})
	Sql = _package.IntelligenceSql(where, Sql, false)
	var data []GetMerchantlst
	var Model []merchant.Merchant
	_package.Paging(Sql, page).Find(&Model)
	for _, s := range Model {
		lst := GetMerchantlst{}
		_package.StructAssign(&lst, &s)
		lst.CreateTime = _package.UnixToDate(s.CreatedAt)
		lst.MerAddress = s.Address + s.DetAddress
		data = append(data, lst)
	}
	err := Sql.Count(&count).Error
	if err != nil {
		return _package.List(data, 0)
	}
	return _package.List(data, count)
}

type PostMerchantCreate struct {
	Pid            int     `json:"p_id"`
	Id             int     `json:"id"`
	IsFlag         int     `json:"is_flag"`
	MerName        string  `json:"mer_name" validate:"required" error:"商户名称"`
	CategoryId     int     `json:"category_id" validate:"required" error:"商户名称"`
	MerAccount     string  `json:"mer_account" validate:"required" error:"商户账号"`
	MerPassword    string  `json:"mer_password"  error:"商户密码"`
	RealName       string  `json:"real_name" validate:"required" error:"商户真实姓名"`
	MerPhone       string  `json:"mer_phone" validate:"required" error:"商户手机号码"`
	CommissionRate float64 `json:"commission_rate"  error:"手续费"`
	MerKeyword     string  `json:"mer_keyword"`
	MerAddress     string  `json:"mer_address"`
	Mark           string  `json:"mark"`
	Sort           int     `json:"sort"`
	Status         int     `json:"status"`
	IsBroRoom      int     `json:"is_bro_room"`
	IsAudit        int     `json:"is_audit"`
	IsBroGoods     int     `json:"is_bro_goods"`
	IsBest         int     `json:"is_best"`
	IsTrader       int     `json:"is_trader"`
}

func FindStoreMerchant(Id int) (merchant.Merchant, error) {
	user := merchant.Merchant{}
	if err := easygorm.GetEasyGormDb().Where("id = ?", Id).First(&user); err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

/** 添加子商户 **/
func (ser *MerchantService) PostCreateSon(param PostMerchantCreate) merchant.Merchant {
	model := merchant.Merchant{}
	Sql := easygorm.GetEasyGormDb().Model(model)
	_package.StructAssign(&model, &param)
	_package.AddStructCommon(0, &model, _package.Authority{})

	if err := Sql.Create(&model); err.Error != nil {
		ser.isErr = err.Error
	}
	return model
}

/** 修改子商户 **/
func (ser *MerchantService) PostUpdateSon(update map[string]interface{}, param PostMerchantCreate) merchant.Merchant {
	model, err := FindStoreMerchant(param.Id)
	if err != nil {
		ser.isErr = err
		return model
	}
	if model.Pid != param.Pid {
		ser.isErr = errors.New("权限不足")
		return model
	}
	_package.StructAssign(&model, &param)
	_package.AddStructCommon(model.Id, &model, _package.Authority{})
	Sql := easygorm.GetEasyGormDb()
	if err := Sql.Updates(&model); err.Error != nil {
		ser.isErr = err.Error
	}
	return model
}

type RegisteredMerchant struct {
	Pid         int    `gorm:"default 0 comment('上级ID') INT(10)"`
	IsFlag      int    `gorm:"not null default 1  TINYINT(1)" json:"is_flag"`
	MerName     string `json:"mer_name"  validate:"required" error:"商户名称"`
	CategoryID  int    `json:"category_id"  validate:"required" error:"行业分类"`
	MerAccount  string `json:"mer_account"  validate:"required" error:"商户账号"`
	MerPassword string `json:"mer_password" validate:"required" error:"商户密码"`
	RealName    string `json:"real_name" validate:"required" error:"商户真实姓名"`
	MerPhone    string `json:"mer_phone" validate:"required" error:"商户手机号码"`
	Code        int    `json:"code" validate:"required" error:"短信验证码"`
	Local       struct {
		Long     string `json:"long" validate:"required"  error:"坐标"`
		Lat      string `json:"lat" validate:"required"  error:"坐标"`
		Province string `json:"province" validate:"required"  error:"省"`
		City     string `json:"city" validate:"required"  error:"城市"`
		CityId   int    `json:"city_id" validate:"required"  error:"百度城市编码"`
		District string `json:"district" validate:"required"  error:"区县"`
		Adcode   int    `json:"adcode" validate:"required"  error:"行政区划代码"`
		Address  string `json:"address" validate:"required"  error:"poi地址"`
	} `json:"local"  validate:"required" error:"选择地图坐标"`
	DetAddress string `json:"det_address" validate:"required"  error:"详细地址"`
}

func registeredInit(param RegisteredMerchant) (merchant.Merchant, error) {
	model := merchant.Merchant{
		Status:     1,
		ExpireType: 1,
	}
	_package.StructAssign(&model, &param)
	model.Long = param.Local.Long
	model.Lat = param.Local.Lat
	model.Province = param.Local.Province
	model.City = param.Local.City
	model.CityId = param.Local.CityId
	model.District = param.Local.District
	model.Adcode = param.Local.Adcode
	model.Address = param.Local.Address
	model.DetAddress = param.DetAddress
	//获取商户试用时间
	sys := new(system.SystemConfigService)
	TrialTime := sys.GetTrialTime()
	if sys.Error() != "" {
		return model, errors.New(sys.Error())
	}
	currentTime := time.Now()
	model.ExpireTime = currentTime.AddDate(0, 0, TrialTime).Unix()
	return model, nil
}

/** pc端注册商户 **/
func (ser *MerchantService) RegisteredMerchant(param RegisteredMerchant) merchant.Merchant {
	model, errs := registeredInit(param)
	if errs != nil {
		ser.isErr = errs
		return model
	}
	mer, err := registeredMerchant(model)
	if err != nil {
		ser.isErr = err
		return mer
	}
	model.Pid = mer.Id
	model.IsFlag = 1
	_, errSon := registeredMerchant(model)
	if errSon != nil {
		ser.isErr = err
		return mer
	}
	return model
}

func registeredMerchant(model merchant.Merchant) (merchant.Merchant, error) {
	Sql := easygorm.GetEasyGormDb().Model(model)
	_package.AddStructCommon(0, &model, _package.Authority{})
	if err := Sql.Create(&model); err.Error != nil {
		return model, err.Error
	}
	return model, nil
}

func (ser *MerchantService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
