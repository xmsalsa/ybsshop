/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-7
 * Time: 下午 18:06
 */
package product

import (
	"shop/application/models"
)

type StoreCategory struct {
	models.Model
	Id         int    `gorm:"not null primary_key INT(10)" json:"id"`
	MerchId    int    `gorm:"not null default 0  INT(11)"  json:"merch_id" json:"merch_id"`
	Pid        int    `gorm:"not null comment('父id') index MEDIUMINT(11)" json:"pid"`
	CateName   string `gorm:"not null comment('分类名称') VARCHAR(100)" json:"cate_name" validate:"required"`
	Sort       int    `gorm:"not null comment('排序') index MEDIUMINT(11)" json:"sort"`
	Pic        string `gorm:"not null default '' comment('图标') VARCHAR(128)" json:"pic"`
	IsShow     int    `gorm:"not null default 1 comment('是否显示') TINYINT(1)" json:"is_show"`
	BigPic     string `gorm:"not null default '' comment('分类大图') VARCHAR(255)" json:"big_pic"`
	CreatedAt  int64  `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid int64  `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt  int64  `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid int64  `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect     int64  `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
