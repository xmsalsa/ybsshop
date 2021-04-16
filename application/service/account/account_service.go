/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 00:45:52
 */
package account

import (
	"shop/application/libs"
	"shop/application/libs/easygorm"
	"shop/application/libs/utils"
	"shop/application/models"

	"gorm.io/gorm"
)

type AccountService struct {
	Gorm *gorm.DB
}

func (s AccountService) getGormDb() *gorm.DB {
	if s.Gorm == nil {
		s.Gorm = easygorm.GetEasyGormDb()
	}
	return s.Gorm
}
func (s AccountService) getGormDbWithModel() *gorm.DB {
	if s.Gorm == nil {
		s.Gorm = easygorm.GetEasyGormDb()
	}
	return s.Gorm.Model(models.Account{})
}

func (s AccountService) AccountExist(username string) (bool, error) {
	var exist bool = false
	account := models.Account{}

	err := s.getGormDbWithModel().Where("username=?", username).Where("effect=1").Find(&account).Error
	if err != nil {
		return exist, err
	}
	if account.ID > 0 {
		exist = true
	}

	return exist, nil
}

// 创建账号
func (s AccountService) AccountCreate(param SAddAccount) (models.Account, error) {
	var err error
	account := models.Account{}

	// 参数验证
	// 略

	// 判断账号重复
	// 略,创建管理员时有判断

	utils.InitModel(&account, param.UpdatedUid, 0)
	utils.StructCopy(&param, &account)
	if account.Password != "" {
		account.Password = libs.HashPassword(account.Password)
	}
	err = s.getGormDbWithModel().Create(&account).Error
	if err != nil {
		return account, nil
	}

	return account, nil
}
