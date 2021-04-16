/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 14:45
 */
package product

import (
	"shop/application/models"
)

type StoreProductRule struct {
	models.CommonModel
	Id         int    `gorm:"not null primary_key INT(10)" json:"id"`
	MerchId    int    `gorm:"not null default 0  INT(11)"  json:"merch_id"`
	RuleName   string `gorm:"not null comment('规格名称') VARCHAR(32)"`
	RuleValue  string `gorm:"not null comment('规格值') TEXT"`
	CreatedAt  int64  `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid int64  `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt  int64  `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid int64  `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect     int64  `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
