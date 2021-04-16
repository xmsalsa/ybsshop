/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-11 23:03:19
 */
package group

import (
	"shop/application/models"
)

type ClientGroup struct {
	models.Model
	Name        string `gorm:"not null; type:VARCHAR(24); default:'';column:name;comment('名称'); " json:"group_name" validate:"required"`
	Pid         int    `gorm:"not null; type:INT; default:0; column:pid; comment('父类ID')" json:"pid"`
	ClientCount int    `gorm:"not null; type:INT; default:0; column:client_count; comment('客户数量')" json:"client_count"`
	MerchId     int    `gorm:"not null; type:INT; default:0; column:merch_id; comment('商户ID')" json:"merch_id"`
	Sort        int    `gorm:"not null; type:INT; default:0; column:sort; comment('排序')" json:"sort"`
}

func (this *ClientGroup) TableComment() string {
	return "用户分组"
}

func (this *ClientGroup) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":           this.Id,
			"group_name":   this.Name,
			"merch_id":     this.MerchId,
			"client_count": this.ClientCount,
			// "sort":       this.Sort,
			// "pid":        this.Pid,
			// "created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
