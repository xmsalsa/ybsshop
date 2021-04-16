package daccount

import (
	"errors"
	"fmt"
	"shop/service/dao/drole"
	"strconv"

	"shop/application/libs"
	"shop/application/libs/easygorm"
	"shop/application/libs/logging"
	"shop/application/models"
)

const ModelName = "用户管理"

type AccountResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Intro     string `json:"introduction"`
	Avatar    string `json:"avatar"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type AccountListResponse struct {
	AccountResponse
	Roles []string `gorm:"-" json:"roles"`
}

type AccountReq struct {
	Name     string `json:"name" `
	Username string `json:"username"`
	Password string `json:"password"`
	Intro    string `json:"introduction"`
	Avatar   string `json:"avatar"`
}

func (u *AccountResponse) ModelName() string {
	return ModelName
}

func Model() *models.Account {
	return &models.Account{}
}

func (u *AccountResponse) All(name, sort, orderBy string, page, pageSize int) (map[string]interface{}, error) {
	var count int64
	var accounts []*AccountListResponse
	db := easygorm.GetEasyGormDb().Model(Model())
	if len(name) > 0 {
		db = db.Where("name", "like", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	if err != nil {
		logging.ErrorLogger.Errorf("get list count err ", err)
		return nil, err
	}

	paginateScope := easygorm.PaginateScope(page, pageSize, sort, orderBy)
	err = db.Scopes(paginateScope).
		Find(&accounts).Error
	if err != nil {
		logging.ErrorLogger.Errorf("get list data err ", err)
		return nil, err
	}
	// 查询用户角色
	getRoles(accounts)

	list := map[string]interface{}{"items": accounts, "total": count, "limit": pageSize}
	return list, nil
}

func getRoles(accounts []*AccountListResponse) {
	var roleIds []string
	accountRoleIds := make(map[uint][]string, 10)
	for _, account := range accounts {
		accountRoleId := easygorm.GetRolesForUser(account.Id)
		accountRoleIds[account.Id] = accountRoleId
		roleIds = append(roleIds, accountRoleId...)
	}

	roles, err := drole.FindInId(roleIds)
	if err != nil {
		logging.ErrorLogger.Errorf("get role get err ", err)
	}

	for _, account := range accounts {
		for _, role := range roles {
			sRoleId := strconv.FormatInt(int64(role.Id), 10)
			if libs.InArrayS(accountRoleIds[account.Id], sRoleId) {
				account.Roles = append(account.Roles, role.Name)
			}
		}
	}
}

func (u *AccountResponse) FindByUserName(username string) error {
	err := easygorm.GetEasyGormDb().Model(Model()).Where("username = ?", username).Find(u).Error
	if err != nil {
		logging.ErrorLogger.Errorf("find user by username ", username, " err ", err)
		return err
	}
	return nil
}

func (u *AccountResponse) Create(object map[string]interface{}) error {
	if username, ok := object["Username"].(string); ok {
		err := u.FindByUserName(username)
		if err != nil {
			logging.ErrorLogger.Errorf("create user find by username get err ", err)
			return err
		}

		if u.Id > 0 {
			return errors.New(fmt.Sprintf("username %s is being used", username))
		}
	}

	err := easygorm.GetEasyGormDb().Model(Model()).Create(object).Error
	if err != nil {
		logging.ErrorLogger.Errorf("create data err ", err)
		return err
	}

	return nil
}

func (u *AccountResponse) Update(id uint, object map[string]interface{}) error {
	err := u.Find(id)
	if err != nil {
		return err
	}
	if u.Username == "username" {
		return errors.New("不能编辑管理员")
	}
	if username, ok := object["Username"].(string); ok {
		err := u.FindByUserName(username)
		if err != nil {
			logging.ErrorLogger.Errorf("create Account find by username get err ", err)
			return err
		}

		if u.Id > 0 && u.Id != id {
			return errors.New(fmt.Sprintf("username %s is being used", username))
		}
	}
	err = easygorm.GetEasyGormDb().Model(Model()).Where("id = ?", id).Updates(object).Error
	if err != nil {
		logging.ErrorLogger.Errorf("update Account  get err ", err)
		return err
	}
	return nil
}

func (u *AccountResponse) Find(id uint) error {
	err := easygorm.GetEasyGormDb().Model(Model()).Where("id = ?", id).Find(u).Error
	if err != nil {
		logging.ErrorLogger.Errorf("find Account err ", err)
		return err
	}
	return nil
}

func (u *AccountResponse) Delete(id uint) error {
	err := easygorm.GetEasyGormDb().Unscoped().Delete(Model(), id).Error
	if err != nil {
		logging.ErrorLogger.Errorf("delete Account by id get  err ", err)
		return err
	}
	return nil
}

// AddRoleForUser add roles for user
func AddRoleForUser(account *models.Account) error {
	if len(account.RoleIds) == 0 {
		return nil
	}

	var err error
	var roleIds []string
	var oldRoleIds []string

	userId := strconv.FormatUint(uint64(account.ID), 10)
	oldRoleIds, err = easygorm.GetEasyGormEnforcer().GetRolesForUser(userId)
	if err != nil {
		logging.ErrorLogger.Errorf("add role to account,del role  err: %+v\n", err)
		return err
	}

	for _, roleId := range account.RoleIds {
		roleId := strconv.FormatUint(uint64(roleId), 10)
		if len(oldRoleIds) > 0 && libs.InArrayS(oldRoleIds, roleId) {
			continue
		}

		roleIds = append(roleIds, roleId)
	}

	if _, err := easygorm.GetEasyGormEnforcer().AddRolesForUser(userId, roleIds); err != nil {
		logging.ErrorLogger.Errorf("add role to user role failed: %+v\n", err)
		return err
	}

	return nil
}

func (u *AccountResponse) Profile(id uint) error {
	return u.Find(id)
}
