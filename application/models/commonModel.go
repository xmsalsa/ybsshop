/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-8
 * Time: 上午 9:36
 */
package models

type CommonModel struct {
	Id         uint  `gorm:"primarykey" json:"id"`
	CreatedAt  int64 `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid int64 `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt  int64 `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid int64 `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect     int64 `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
