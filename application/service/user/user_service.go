/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 09:32:13
 */
package user

import (
	"errors"
	"fmt"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	"shop/application/models"
)

type UserDetail struct {
	Uid        int    `form:"uid" json:"uid" validate:"omitempty,gt=0" comment:"用户ID"`
	Phone      string `form:"phone" validate:"omitempty" json:"phone"`
	RealName   string `form:"real_name" json:"real_name"`
	Nickname   string `form:"nickname" json:"nickname"`
	Birthday   int64  `form:"birthday" json:"birthday"`
	CardId     string `form:"card_id" json:"card_id"`
	Mark       string `form:"mark" json:"mark"`
	Avatar     string `form:"avatar" json:"avatar"`
	Address    string `form:"address" json:"address"`
	UpdatedUid int    `form:"updated_uid" json:"updated_uid"`
}
type UserCreate struct {
	Phone      string `form:"phone" validate:"omitempty" json:"phone"`
	RealName   string `form:"real_name" json:"real_name"`
	Nickname   string `form:"nickname" json:"nickname"`
	Birthday   int64  `form:"birthday" json:"birthday"`
	CardId     string `form:"card_id" json:"card_id"`
	Mark       string `form:"mark" json:"mark"`
	Avatar     string `form:"avatar" json:"avatar"`
	Address    string `form:"address" json:"address"`
	UpdatedUid int    `form:"updated_uid" json:"updated_uid"`
}

func repsToDesc(user models.User) UserDetail {
	return UserDetail{
		Uid:   user.Id,
		Phone: user.Phone,
	}
}

// 通过手机号查找客户时, 不存则创建
func GetUser(param UserDetail) (models.User, error) {
	u := models.User{}

	// 验证手机号码是否正确
	if param.Phone != "" {
		// 略
	}

	err := utils.Validate.Struct(param)
	if err != nil {
		return u, err
	}

	condition := make(map[string]interface{})
	if param.Uid > 0 {
		condition["id"] = param.Uid
	}
	if param.Phone != "" {
		condition["phone"] = param.Phone
	}
	if len(condition) > 0 {
		// 查找
		db := utils.GetGormDbWithModel(u)
		utils.Build(db, condition)
		err = db.Find(&u).Error
		if err != nil {
			return u, nil
		}
	}

	if u.Id == 0 {
		if param.Uid > 0 {
			return u, errors.New(fmt.Sprintf("%sID:%s%s", u.TableComment(), param.Uid, response.DB_RECORD_NOEXSIT))
		}
		adduser_param := UserCreate{}
		utils.StructCopy(&param, &adduser_param)
		u, err = create(adduser_param)
		if err != nil {
			return u, nil
		}
	}

	return u, nil
}

func create(param UserCreate) (models.User, error) {
	user := models.User{}
	utils.InitModel(&user, param.UpdatedUid, 0)
	utils.StructCopy(&param, &user)

	err := utils.GetGormDbWithModel(models.User{}).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
