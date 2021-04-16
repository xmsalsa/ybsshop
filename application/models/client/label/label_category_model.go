/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-07 18:02:06
 */
package label

import (
	"shop/application/libs/utils"
	"shop/application/models"
)

//定义标签分类结构体
type LabelCategory struct {
	models.Model
	MerchId int    `gorm:"not null; type:INT; default:0; column:merch_id; comment('商户ID')" json:"merch_id"`
	Name    string `gorm:"not null; type:VARCHAR(24); default:'';column:name;comment('名称'); " json:"labelcate_name" validate:"required"`
	Pid     int    `gorm:"not null; type:INT; default:0; column:pid; comment('父类ID')" json:"pid"`
	Type    int    `gorm:"not null; type:INT; default:1; column:type; comment('1=用户;')" json:"type"`
	Sort    int    `gorm:"not null; type:INT; default:0; column:sort; comment('排序')" json:"sort"`
	OwnerId int    `gorm:"not null; type:INT; default:0; column:owner_id;  comment('所有人, 0为全部')" json:"owner_id"`
	Other   string `gorm:"not null; type:text; column:other; comment('名称');" json:"other"`
}

func (this *LabelCategory) TableComment() string {
	return "标签分类"
}

func (this *LabelCategory) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":             this.Id,
			"labelcate_name": this.Name,
			"merch_id":       this.MerchId,
			"sort":           this.Sort,
			// "pid":            this.Pid,
			// "type":           this.Type,
			// "owner_id":       this.OwnerId,
			// "other":          this.Other,
			"created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}
