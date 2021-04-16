/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 11:23:28
 */
package label

import (
	"shop/application/models"
)

type ClientLabelRelation struct {
	models.Model
	MerchId  int `gorm:"not null; type:INT; default:0; column:merch_id; comment('商户ID')" json:"merch_id"`
	ClientId int `gorm:"not null; type:INT; default:0; column:client_id; comment('客户ID')" json:"client_id"`
	LabelId  int `gorm:"not null; type:INT; default:0; column:label_id; comment('标签ID')" json:"label_id"`
}

func (this *ClientLabelRelation) TableComment() string {
	return "客户标签关联"
}

func (this *ClientLabelRelation) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":        this.Id,
			"merch_id":  this.MerchId,
			"client_id": this.ClientId,
			"group_id":  this.LabelId,
			// "created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
