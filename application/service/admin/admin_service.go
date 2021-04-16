/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 00:45:38
 */
package admin

import (
	"errors"
	"fmt"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	"shop/application/models"
	madmin "shop/application/models/admin"
	saccount "shop/application/service/account"
	suser "shop/application/service/user"
)

// 创建管理员
func AdminCreate(param SAdminCreate) (map[string]interface{}, error) {
	var err error
	data := make(map[string]interface{})
	user := models.User{}
	admins := madmin.Admin{}
	account := models.Account{}

	// 前提, 账号不存在
	var exist bool = true
	exist, err = saccount.AccountExist(param.Username)
	if err != nil {
		return data, err
	}
	if exist {
		return data, errors.New(fmt.Sprintf("%s%s%s", account.TableComment(), param.Username, response.DB_RECORD_EXSIT))
	}

	// 获取用户, 拿Uid
	user_param := suser.UserDetail{}
	utils.StructCopy(&param, &user_param)
	user, err = suser.GetUser(user_param)
	if err != nil {
		return data, err
	}

	// 创建管理员, 拿Id
	addAdmin_param := SAddAdmin{}
	utils.StructCopy(&param, &addAdmin_param)
	addAdmin_param.Uid = user.Id
	admins, err = AddAdmin(addAdmin_param)
	if err != nil {
		return data, err
	}

	// 创建账号
	account_param := saccount.SAddAccount{}
	utils.StructCopy(&param, &account_param)
	account_param.Identity = admins.Id // 管理员ID
	account_param.Type = 1             //1 管理员, 2客户
	account_param.Status = 1           //1 正常
	account, err = saccount.AccountCreate(account_param)
	if err != nil {
		return data, err
	}

	data["admin"] = map[string]interface{}{
		"admin_id":  admins.Id,
		"uid":       admins.Uid,
		"real_name": admins.RealName,
		"nickname":  admins.Nickname,
		"phone":     admins.Phone,
	}
	data["account"] = map[string]interface{}{
		"account_id": account.ID,
		"username":   account.Username,
	}
	return data, nil
}

// 创建管理员
func AddAdmin(param SAddAdmin) (madmin.Admin, error) {
	var err error
	admins := madmin.Admin{}

	// 参数验证
	// 略, 该方法会开放给外总调用, 需要判断参数

	utils.InitModel(&admins, param.UpdatedUid, param.MerchId)
	utils.StructCopy(&param, &admins)
	err = utils.GetGormDbWithModel(admins).Create(&admins).Error
	if err != nil {
		return admins, nil
	}

	return admins, err
}
