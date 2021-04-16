/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-10 14:52:21
 */
package label

import (
	"shop/application/models"
)

//定义标签分类结构体
type Label struct {
	models.Model
	MerchId     int    `gorm:"not null; type:INT; default:0; column:merch_id; comment('商户ID')" json:"merch_id"`
	Name        string `gorm:"not null; type:VARCHAR(24); default:'';column:name;comment('名称'); " json:"label_name" validate:"required"`
	LabelcateId int    `gorm:"not null; type:INT; default:0; column:labelcate_id; comment('标签分类ID')" json:"labelcate_id"`
	ClientCount int    `gorm:"not null; type:INT; default:0; column:client_count; comment('客户数量')" json:"client_count"`
	Type        int    `gorm:"not null; type:tinyint(1); default:0; column:type; comment('1客户标签')" json:"type"`
	Source      int    `gorm:"not null; type:tinyint(1); default:0; column:source; comment('1=手动标签 2=自动标签')" json:"source"`
	Sort        int    `gorm:"not null; type:INT; default:0; column:sort; comment('排序DESC')" json:"sort"`
}

func (this *Label) TableComment() string {
	return "标签"
}
func (this *Label) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":           this.Id,
			"label_name":   this.Name,
			"labelcate_id": this.LabelcateId,
			"client_count": this.ClientCount,
			// "merch_id":   this.MerchId,
			// "created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
