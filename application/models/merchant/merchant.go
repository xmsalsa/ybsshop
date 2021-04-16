/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 14:24
 */
package merchant

type Merchant struct {
	Id             int     `gorm:"not null primary_key INT(10)" json:"id"`
	Pid            int     `gorm:"default 0 comment('上级ID') INT(10)"`
	IsFlag         int     `gorm:"not null default 1  TINYINT(1)" json:"is_flag"`
	CategoryId     int     `gorm:"not null default 0 comment('商户分类 id') INT(10)"`
	MerName        string  `gorm:"not null comment('商户名称') VARCHAR(32)"`
	RealName       string  `gorm:"not null comment('商户姓名') VARCHAR(32)"`
	MerPhone       string  `gorm:"not null comment('商户手机号') VARCHAR(13)"`
	MerAddress     string  `gorm:"not null comment('商户地址') VARCHAR(64)"`
	MerKeyword     string  `gorm:"not null comment('商户关键字') VARCHAR(64)"`
	MerAvatar      string  `gorm:"comment('商户头像') VARCHAR(128)"`
	MerBanner      string  `gorm:"comment('商户banner图片') VARCHAR(128)"`
	MiniBanner     string  `gorm:"comment('商户店店铺街图片') VARCHAR(128)"`
	Sales          int     `gorm:"default 0 comment('销量') INT(11)"`
	ProductScore   float64 `gorm:"default 5.0 comment('商品描述评分') DECIMAL(11,1)"`
	ServiceScore   float64 `gorm:"default 5.0 comment('服务评分') DECIMAL(11,1)"`
	PostageScore   float64 `gorm:"default 5.0 comment('物流评分') DECIMAL(11,1)"`
	Mark           string  `gorm:"not null comment('商户备注') VARCHAR(256)"`
	RegAdminId     int     `gorm:"not null default 0 comment('总后台管理员ID') INT(10)"`
	Sort           int     `gorm:"not null default 0 INT(10)"`
	ExpireType     int     `gorm:"not null default 1 comment('1试用0永久2限制时间') TINYINT(1)"`
	ExpireTime     int64   `gorm:"not null default 1 comment('过期时间') int(10)"`
	Status         int     `gorm:"not null default 0 comment('商户是否禁用0锁定,1正常') TINYINT(1)"`
	CommissionRate float64 `gorm:"comment('提成比例') DECIMAL(11,4)"`
	Long           string  `gorm:"comment('经度') VARCHAR(16)"`
	Lat            string  `gorm:"comment('纬度') VARCHAR(16)"`
	Province       string  `gorm:"comment('省') VARCHAR(55)"`
	City           string  `gorm:"comment('城市') VARCHAR(55)"`
	CityId         int     `gorm:"comment('百度城市编码') int(10)"`
	District       string  `gorm:"comment('区县') VARCHAR(255)"`
	Adcode         int     `gorm:"comment('行政区划代码') INT(10)"`
	Address        string  `gorm:"comment('poi地址') VARCHAR(255)"`
	DetAddress     string  `gorm:"comment('详细地址') VARCHAR(255)"`
	//	IsDel          int       `gorm:"not null default 0 comment('0未删除1删除') TINYINT(1)"`
	IsAudit        int    `gorm:"not null default 0 comment('添加的产品是否审核0不审核1审核') TINYINT(1)"`
	IsBroRoom      int    `gorm:"not null default 1 comment('是否审核直播间0不审核1审核') TINYINT(1)"`
	IsBroGoods     int    `gorm:"not null default 1 comment('是否审核直播商品0不审核1审核') TINYINT(1)"`
	IsBest         int    `gorm:"not null default 0 comment('是否推荐') TINYINT(1)"`
	IsTrader       int    `gorm:"not null default 0 comment('是否自营') TINYINT(1)"`
	MerState       int    `gorm:"not null default 0 comment('商户是否1开启0关闭') TINYINT(1)"`
	MerInfo        string `gorm:"not null default '' comment('店铺简介') VARCHAR(256)"`
	ServicePhone   string `gorm:"not null default '' comment('店铺电话') VARCHAR(13)"`
	CareCount      int    `gorm:"default 0 comment('关注总数') INT(11)"`
	CopyProductNum int    `gorm:"default 0 comment('剩余复制商品次数') INT(11)"`
	CreatedAt      int64  `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid     int64  `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt      int64  `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid     int64  `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect         int64  `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
