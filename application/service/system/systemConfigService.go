/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-14
 * Time: 上午 10:58
 */
package system

import (
	"encoding/json"
	"shop/application/libs/easygorm"
	system2 "shop/application/libs/system"
	"shop/application/models/system"
)

/* 定义结构体 */
type SystemConfigService struct {
	/* 错误体 */
	isErr error
}

func (ser *SystemConfigService) GetSystemConfig(keys string) string {
	system := system.SystemConfig{}
	if err := easygorm.GetEasyGormDb().Where("k = ?", keys).First(&system); err.Error != nil {
		ser.isErr = err.Error
		return ""
	}
	return system.Config
}

type TrialTime struct {
	TrialTime int `json:"trial_time"`
}

func (ser *SystemConfigService) GetTrialTime() int {
	res := 0
	config := ser.GetSystemConfig(system2.TRIAL_TIME)
	if config == "" {
		return res
	}
	TrialTime := TrialTime{}
	err := json.Unmarshal([]byte(config), &TrialTime)
	if err != nil {
		ser.isErr = err
		return res
	}
	return TrialTime.TrialTime
}

func (ser *SystemConfigService) Error() string {
	if ser.isErr != nil {
		return ser.isErr.Error()
	}
	return ""
}
