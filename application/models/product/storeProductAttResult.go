/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-9
 * Time: 下午 17:44
 */
package product

type StoreProductAttrsResult struct {
	ProductId int `gorm:"not null comment('商品ID') index INT(10)"`
	//MerchId        int    `gorm:"not null default 0 comment('商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)') INT(10)"`
	Result     string `gorm:"not null comment('商品属性参数') LONGTEXT"`
	ChangeTime int64  `gorm:"not null comment('上次修改时间') INT(10)"`
	Type       int    `gorm:"default 0 comment('活动类型 0=商品，1=秒杀，2=砍价，3=拼团') TINYINT(1)"`
}
