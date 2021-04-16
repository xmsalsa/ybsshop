/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 00:40:58
 */
package admin

import (
	"shop/application/libs/utils"
	"shop/application/models"
)

func (this *Admin) TableComment() string {
	return "管理员"
}

type Admin struct {
	models.Model

	MerchId  int    `gorm:"not null; type:INT; default:0; column:merch_id; comment('商户ID')" json:"merch_id"`
	Uid      int    `gorm:"not null; type:INT; default:0; column:uid; comment('用户ID')" json:"uid"`
	RealName string `gorm:"not null;type:varchar(32);column:real_name;comment('真实姓名');" json:"real_name"`
	Nickname string `gorm:"not null;type:varchar(32);column:nickname;comment('昵称');" json:"nickname"`
	Phone    string `gorm:"type:VARCHAR(15); default:'';column:phone;comment('手机号码'); " json:"phone"`
	Email    string `gorm:"type:VARCHAR(128); default:'';column:email;comment('邮箱'); " json:"email"`
	Avatar   string `gorm:"type:varchar(255)"column:avatar;comment('头像'); json:"avatar"`
	UseApp   int    `gorm:"type:uint;default:0;column:use_app;comment('使用APP:1是;0否');" json:"use_app"`
	UsePc    int    `gorm:"type:uint;default:0;column:use_pc;comment('使用PC:1是;0否');" json:"use_pc"`
	Type     int    `gorm:"type:uint;default:2;comment('账号类型:1系统管理员;2商户管理员');" json:"type"`
	Status   int    `gorm:"type:INT; default:1; column:status; comment('1正常，0离职,')" json:"status"`
	LastIp   string `gorm:"type:varchar(24);default:'';column:last_ip" json:"last_ip" form:"last_ip"`
	LastTime int64  `gorm:"type:INT;default:0;column:last_time" json:"last_time" form:"last_time"`
	// 其他字段
	// 略
}

func (this *Admin) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":         this.Id,
			"merch_id":   this.MerchId,
			"uid":        this.Uid,
			"real_name":  this.RealName,
			"nickname":   this.Nickname,
			"phone":      this.Phone,
			"email":      this.Email,
			"avatar":     this.Avatar,
			"use_app":    this.UseApp,
			"use_pc":     this.UsePc,
			"type":       this.Type,
			"status":     this.Status,
			"last_ip":    this.LastIp,
			"last_time":  this.LastTime,
			"created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
