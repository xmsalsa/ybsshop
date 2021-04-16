/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-07 16:01:04
 */
package models

type Model struct {
	Id         int    `gorm:"not null; primary_key;AUTO_INCREMENT;type:int;column:id;" json:"id"` //主键 自增
	CreatedAt  int64  `gorm:"not null; type:INT; default:0; column:created_at; comment('创建时间')" json:"created_at"`
	CreatedUid int    `gorm:"not null; type:INT; default:0; column:created_uid; comment('创建人ID')" json:"created_uid"`
	UpdatedAt  int64  `gorm:"not null; type:INT; default:0; column:updated_at; comment('更新时间')" json:"updated_at"`
	UpdatedUid int    `gorm:"not null; type:INT; default:0; column:updated_uid; comment('创建人ID')" json:"updated_uid"`
	Effect     int    `gorm:"not null; type:tinyint(1); default:0; column:effect; comment('1有效;0逻辑删')" json:"effect"`
	Memo       string `gorm:"not null; ttype:VARCHAR(255);  default:''; column:memo; comment('备注')" json:"memo"`
}
