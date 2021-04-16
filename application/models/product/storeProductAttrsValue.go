/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-11
 * Time: 下午 13:18
 */
package product

type StoreProductAttrsValue struct {
	ProductId    int     `gorm:"not null comment('商品ID') index(store_id) INT(10)"`
	MerchId      int     `gorm:"not null default 0 comment('商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)') INT(10)"`
	Suk          string  `gorm:"not null comment('商品属性索引值 (attr_value|attr_value[|....])') index(store_id) index(unique) VARCHAR(128)"`
	Stock        int     `gorm:"not null comment('属性对应的库存') INT(10)"`
	Sales        int     `gorm:"not null default 0 comment('销量') INT(10)"`
	Price        int     `gorm:"not null comment('属性金额') int(10)"`
	Image        string  `gorm:"comment('图片') VARCHAR(128)"`
	Unique       string  `gorm:"not null default '' comment('唯一值') index(unique) VARCHAR(32)"`
	Cost         int     `gorm:"not null comment('成本价') int(10)"`
	BarCode      string  `gorm:"not null default '' comment('商品条码') VARCHAR(50)"`
	OtPrice      int     `gorm:"not null default 0.00 comment('原价') DECIMAL(10)"`
	VipPrice     int     `gorm:"not null default 0.00 comment('会员专享价') DECIMAL(10)"`
	Weight       float32 `gorm:"not null default 0.00 comment('重量') DECIMAL(8,2)"`
	Volume       float32 `gorm:"not null default 0.00 comment('体积') DECIMAL(8,2)"`
	Brokerage    int     `gorm:"not null default 0.00 comment('一级返佣') int(10)"`
	BrokerageTwo int     `gorm:"not null default 0.00 comment('二级返佣') int(10)"`
	Type         int     `gorm:"default 0 comment('活动类型 0=商品，1=秒杀，2=砍价，3=拼团') TINYINT(1)"`
	Quota        int     `gorm:"comment('活动限购数量') INT(11)"`
	QuotaShow    int     `gorm:"comment('活动限购数量显示') INT(11)"`
	CreatedAt    int64   `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid   int     `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt    int64   `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid   int     `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect       int     `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
