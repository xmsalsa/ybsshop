/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 22:42:14
 */
package models

import (
	"shop/application/libs/utils"
)

func (this *User) TableComment() string {
	return "用户"
}

// User struct is a row record of the eb_user table in the employees database
type User struct {
	Model

	RealName string `gorm:"type:VARCHAR(32); default:'';column:real_name;comment('真实姓名'); " json:"real_name"`
	Nickname string `gorm:"type:VARCHAR(32); default:'';column:nickname;comment('昵称'); " json:"nickname"`
	Phone    string `gorm:"type:VARCHAR(15); default:'';column:phone;comment('手机号码'); " json:"phone"`
	Birthday int64  `gorm:"type:INT; default:0; column:birthday; comment('生日')" json:"birthday"`
	CardId   string `gorm:"type:VARCHAR(24); default:'';column:card_id;comment('身份证号码'); " json:"card_id"`
	Mark     string `gorm:"type:VARCHAR(255); default:'';column:mark;comment('客户备注'); " json:"mark"`
	Avatar   string `gorm:"type:VARCHAR(255); default:'';column:avatar;comment('头像'); " json:"avatar"`
	Address  string `gorm:"type:VARCHAR(255); default:'';column:address;comment('详细地址'); " json:"address"`
}

func (this *User) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"uid": 0}
	} else {
		return map[string]interface{}{
			"uid":        this.Id,
			"real_name":  this.RealName,
			"nickname":   this.Nickname,
			"phone":      this.Phone,
			"created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
