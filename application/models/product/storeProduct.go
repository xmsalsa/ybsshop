/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-8
 * Time: 下午 14:51
 */
package product

type StoreProduct struct {
	Id             int    `gorm:"not null pk autoincr comment('商品id') MEDIUMINT(11)"`
	MerchId        int    `gorm:"not null default 0 comment('商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)') INT(10)"`
	Image          string `gorm:"not null comment('商品图片') VARCHAR(256)"`
	RecommendImage string `gorm:"not null default '' comment('推荐图') VARCHAR(256)"`
	SliderImage    string `gorm:"not null comment('轮播图') VARCHAR(2000)"`
	StoreName      string `gorm:"not null comment('商品名称') VARCHAR(128)"`
	StoreInfo      string `gorm:"not null comment('商品简介') VARCHAR(256)"`
	Keyword        string `gorm:"not null comment('关键字') VARCHAR(256)"`
	BarCode        string `gorm:"not null default '' comment('商品条码（一维码）') VARCHAR(15)"`
	CateId         string `gorm:"not null comment('分类id') index VARCHAR(64)"`
	Price          int    `gorm:"not null default 0.00 comment('商品价格') index int(10)"`
	VipPrice       int    `gorm:"not null default 0.00 comment('会员价格') int(10)"`
	OtPrice        int    `gorm:"not null default 0.00 comment('市场价') int(10)"`
	Postage        int    `gorm:"not null default 0.00 comment('邮费') int(10)"`
	UnitName       string `gorm:"not null comment('单位名') VARCHAR(32)"`
	Sort           int    `gorm:"not null default 0 comment('排序') index SMALLINT(11)"`
	Sales          int    `gorm:"not null default 0 comment('销量') index MEDIUMINT(11)"`
	Stock          int    `gorm:"not null default 0 comment('库存') MEDIUMINT(11)"`
	IsShow         int    `gorm:"not null default 1 comment('状态（0：未上架，1：上架）') index TINYINT(1)"`
	IsHot          int    `gorm:"not null default 0 comment('是否热卖') index TINYINT(1)"`
	IsBenefit      int    `gorm:"not null default 0 comment('是否优惠') index TINYINT(1)"`
	IsBest         int    `gorm:"not null default 0 comment('是否精品') index TINYINT(1)"`
	IsNew          int    `gorm:"not null default 0 comment('是否新品') index TINYINT(1)"`
	IsPostage      int    `gorm:"not null default 0 comment('是否包邮') index TINYINT(1)"`
	MerUse         int    `gorm:"not null default 0 comment('商户是否代理 0不可代理1可代理') TINYINT(1)"`
	GiveIntegral   int    `gorm:"not null comment('获得积分') int(10)"`
	Cost           int    `gorm:"not null comment('成本价') int(10)"`
	IsSeckill      int    `gorm:"not null default 0 comment('秒杀状态 0 未开启 1已开启') TINYINT(1)"`
	IsBargain      int    `gorm:"comment('砍价状态 0未开启 1开启') TINYINT(1)"`
	IsGood         int    `gorm:"not null default 0 comment('是否优品推荐') TINYINT(1)"`
	IsSub          int    `gorm:"not null default 0 comment('是否单独分佣') TINYINT(1)"`
	IsVip          int    `gorm:"not null default 0 comment('是否开启会员价格') TINYINT(1)"`
	Ficti          int    `gorm:"default 100 comment('虚拟销量') MEDIUMINT(11)"`
	Browse         int    `gorm:"default 0 comment('浏览量') INT(11)"`
	CodePath       string `gorm:"not null default '' comment('商品二维码地址(用户小程序海报)') VARCHAR(64)"`
	SoureLink      string `gorm:"default '' comment('淘宝京东1688类型') VARCHAR(255)"`
	VideoLink      string `gorm:"not null default '' comment('主图视频链接') VARCHAR(500)"`
	TempId         int    `gorm:"not null default 1 comment('运费模板ID') INT(11)"`
	SpecType       int    `gorm:"not null default 0 comment('规格 0单 1多') TINYINT(1)"`
	Activity       string `gorm:"not null default '' comment('活动显示排序1=秒杀，2=砍价，3=拼团') VARCHAR(255)"`
	Spu            string `gorm:"not null default '' comment('商品SPU') CHAR(13)"`
	LabelId        string `gorm:"not null default '' comment('标签ID') VARCHAR(64)"`
	CommandWord    string `gorm:"not null default '' comment('复制口令') VARCHAR(255)"`
	CreatedAt      int64  `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid     int    `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt      int64  `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid     int    `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect         int    `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
