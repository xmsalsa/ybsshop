/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 09:17:16
 */
package models

// EbSystemAdmin struct is a row record of the eb_system_admin table in the employees database
type EbSystemAdmin struct {
	Id         int64  `gorm:"column:id" json:"id" form:"id"` // 后台管理员表ID
	MerchId    int64  `gorm:"column:mer_id" json:"mer_id" form:"mer_id"`
	Phone      string `gorm:"column:phone" json:"phone" form:"phone"`
	RealName   string `gorm:"column:real_name;type:varchar;" json:"real_name"` // 后台管理员姓名
	Intro      string `gorm:"not null; type:varchar(512)" json:"introduction"`
	Avatar     string `gorm:"type:varchar(1024)" json:"avatar"`
	UseApp     int64  `gorm:"column:use_app" json:"use_app" form:"use_pc"`
	UsePc      int64  `gorm:"column:use_pc" json:"use_pc" form:"use_pc"`
	LastIp     string `gorm:"column:last_ip" json:"last_ip" form:"last_ip"`
	LastTime   int64  `gorm:"column:last_time" json:"last_time" form:"last_time"`
	AddTime    uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"`       // 后台管理员添加时间
	LoginCount uint32 `gorm:"column:login_count;type:uint;default:0;" json:"login_count"` // 登录次数
	Level      uint32 `gorm:"column:level;type:utinyint;default:1;" json:"level"`         // 后台管理员级别
	Status     uint32 `gorm:"column:status;type:utinyint;default:1;" json:"status"`       // 后台管理员状态 1有效0无效
	AdminType  int64  `gorm:"column:admin_type" json:"admin_type" form:"admin_type"`
}
