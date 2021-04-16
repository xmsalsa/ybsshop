/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-14
 * Time: 上午 10:54
 */
package system

type SystemConfig struct {
	Name   string `gorm:"not null comment('系统中文标识') VARCHAR(255)" json:"name" `
	K      string `gorm:"not null comment('系统标识') VARCHAR(255)" json:"k"`
	Config string `gorm:"not null default '' comment('json配置') " json:"config"`
}
