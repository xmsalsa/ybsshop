/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-12
 * Time: 上午 10:39
 */
package product

type StoreProductCate struct {
	Id         int   `xorm:"not null pk autoincr INT(11)"`
	MerchId    int   `gorm:"not null default 0  INT(11)"  json:"merch_id"`
	ProductId  int   `xorm:"not null default 0 comment('商品id') INT(11)"`
	CateId     int   `xorm:"not null default 0 comment('分类id') INT(11)"`
	CatePid    int   `xorm:"not null default 0 comment('一级分类id') INT(11)"`
	Status     int   `xorm:"not null default 0 comment('商品状态') TINYINT(1)"`
	CreatedAt  int64 `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid int64 `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt  int64 `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid int64 `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect     int64 `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
