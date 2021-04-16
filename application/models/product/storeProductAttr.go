/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-9
 * Time: 下午 15:04
 */
package product

type StoreProductAttr struct {
	Id         int    `gorm:"not null pk autoincr INT(10)"`
	ProductId  int    `gorm:"not null default 0 comment('商品ID') index INT(10)"`
	AttrName   string `gorm:"not null comment('属性名') VARCHAR(32)"`
	AttrValues string `gorm:"not null comment('属性值') LONGTEXT"`
	Type       int    `gorm:"default 0 comment('活动类型 0=商品，1=秒杀，2=砍价，3=拼团') TINYINT(1)"`
}
