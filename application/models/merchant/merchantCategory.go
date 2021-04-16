/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-13
 * Time: 下午 14:27
 */
package merchant

type MerchantsCategory struct {
	Id             int    `gorm:"not null pk autoincr comment('商户分类 id') INT(10)"`
	CommissionRate string `gorm:"not null default 0.0000 comment('手续费') DECIMAL(6,4)"`
	CategoryName   string `gorm:"not null comment('商户分类名称') VARCHAR(32)"`
	Config         string `gorm:"not null comment('商户分类名称')"`
	CreatedAt      int64  `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid     int64  `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt      int64  `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid     int64  `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect         int64  `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
