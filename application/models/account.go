/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 09:17:16
 */
package models

import "shop/application/libs/utils"

func (this *Account) TableComment() string {
	return "账号"
}

type Account struct {
	ID       int    `gorm:"not null; primary_key;AUTO_INCREMENT;type:int;column:id;" json:"id"` //主键 自增
	Username string `gorm:"not null;type:varchar(32);column:username;comment('账号');" json:"username"`
	Password string `gorm:"type:varchar(255)"column:password;comment('密码'); json:"password"`
	Identity int    `gorm:"column:info_id;type:uint;default:0;column:identity;comment('账号关联的ID')" json:"identity"` // 客户表client_id或管理员表admin_id
	Type     int    `gorm:"column:type;type:uint;default:1;comment('账号类型:1系统管理员;2商户管理员;3客户');" json:"type"`        // 是否为管理员
	Status   int    `gorm:"type:INT; default:1; column:status; comment('1正常，禁止,2注销')" json:"status"`
	Phone    string `gorm:"type:VARCHAR(15); default:'';column:phone;comment('手机号码'); " json:"phone"`
	Email    string `gorm:"type:VARCHAR(128); default:'';column:email;comment('邮箱'); " json:"email"`

	CreatedAt  int64  `gorm:"not null; type:INT; default:0; column:created_at; comment('创建时间')" json:"created_at"`
	CreatedUid int    `gorm:"not null; type:INT; default:0; column:created_uid; comment('创建人ID')" json:"created_uid"`
	UpdatedAt  int64  `gorm:"not null; type:INT; default:0; column:updated_at; comment('更新时间')" json:"updated_at"`
	UpdatedUid int    `gorm:"not null; type:INT; default:0; column:updated_uid; comment('创建人ID')" json:"updated_uid"`
	Effect     int    `gorm:"not null; type:tinyint(1); default:0; column:effect; comment('1有效;0逻辑删')" json:"effect"`
	Memo       string `gorm:"not null; ttype:VARCHAR(255);  default:''; column:memo; comment('备注')" json:"memo"`

	Oplogs  []Oplog
	RoleIds []uint `gorm:"-" json:"role_ids"`
}

func (this *Account) RepsToDesc() map[string]interface{} {
	if this.ID == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":         this.ID,
			"username":   this.Username,
			"identity":   this.Identity,
			"type":       this.Type,
			"status":     this.Status,
			"phone":      this.Phone,
			"email":      this.Email,
			"created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
