package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"shop/service/dao/daccount"
	"time"

	"shop/application/libs/easygorm"
	"shop/application/models"
	madmin "shop/application/models/admin"
	"shop/service/dao/drole"

	"github.com/jinzhu/configor"

	//"shop/service/dao/duser"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"shop/application/libs"

	"github.com/azumads/faker"
)

var Fake *faker.Faker

var Seeds = struct {
	Perms []struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayname"`
		Description string `json:"description"`
		Act         string `json:"act"`
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())
}

var config = flag.String("config", "", "配置路径")
var path = flag.String("path", "", "数据路径")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] [command]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  --config <path>\n")
		fmt.Fprintf(os.Stderr, "    设置配置文件路径\n")
		fmt.Fprintf(os.Stderr, "  --path <path>\n")
		fmt.Fprintf(os.Stderr, "    设置填充数据路径\n")
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.Parse()

	fpaths, err := filepath.Glob(filepath.Join(*path, "*.yml"))
	if err != nil {
		panic(fmt.Sprintf("数据填充YML文件路径加载失败: %+v\n", err))
	}

	fmt.Printf("数据填充YML文件路径：%+v\n", fpaths)

	if err := configor.Load(&Seeds, fpaths...); err != nil {
		panic(fmt.Sprintf("load config file err：%+v", err))
	}
	err = libs.InitConfig(*config)
	if err != nil {
		panic(fmt.Sprintf("系统配置初始化失败: %+v\n", err))
	}

	err = easygorm.Init(&easygorm.Config{
		Adapter: libs.Config.DB.Adapter,
		Conn:    libs.Config.DB.Conn,
		GormConfig: &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: libs.Config.DB.Prefix,
			},
		},
		Casbin: &easygorm.Casbin{
			Path:   libs.Config.Casbin.Path,
			Prefix: libs.Config.Casbin.Prefix,
		},
		Models: []interface{}{
			&models.Account{},
			&models.User{},
			&madmin.Admin{},
			&models.Role{},
			&models.Permission{},
			&models.Config{},
			&models.Oplog{},
		},
	})
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败: %+v\n", err))
	}

	Seed()

}

func Seed() {
	AutoMigrates()
	perms := CreatePerms()
	CreateRole(perms)
	CreateAdmin()
	//CreateConfigs()
}

// CreateConfigs 新建权限
//func CreateConfigs() {
//	configs := []*models.Config{
//		{
//			Name:  "imageHost",
//			Value: "https://www.snowlyg.com",
//		},
//		{
//			Name:  "beianhao",
//			Value: "",
//		},
//	}
//	logging.DebugLogger.Debugf("系统设置填充：%+v\n", configs)
//		perm, err := models.GetConfig(s)
//		if err != nil && models.IsNotFound(err) {
//			if perm.ID == 0 {
//				perm = &models.Config{
//					Model: gorm.Model{CreatedAt: time.Now()},
//					Name:  m.Name,
//					Value: m.Value,
//				}
//				if err := perm.CreateConfig(); err != nil {
//					logging.ErrorLogger.Errorf("seeder data create config err：%+v\n", err)
//					return
//				}
//			}
//		}
//	}
//}

//CreatePerms 新建权限
func CreatePerms() [][]string {
	var insertPerms []models.Permission
	for _, m := range Seeds.Perms {
		perm := models.Permission{
			Name:        m.Name,
			DisplayName: m.DisplayName,
			Description: m.Description,
			Act:         m.Act,
		}
		insertPerms = append(insertPerms, perm)
	}

	if len(insertPerms) == 0 {
		return nil
	}

	create := easygorm.GetEasyGormDb().Create(&insertPerms)
	if err := create.Error; err != nil {
		panic(fmt.Sprintf("seeder data create perms err：%+v\n", err))
		return nil
	}

	fmt.Println(fmt.Sprintf("\n填充权限数据："))
	for _, insertPerm := range insertPerms {
		fmt.Println(fmt.Sprintf("  %s:%s", insertPerm.Name, insertPerm.Act))
	}

	var perms [][]string
	for _, perm := range insertPerms {
		perms = append(perms, []string{
			perm.Name,
			perm.Act,
		})
	}

	return perms
}

// CreateRole 新建管理角色
func CreateRole(perms [][]string) {
	role := &models.Role{
		Name:        libs.Config.Admin.Rolename,
		DisplayName: libs.Config.Admin.Rolename,
		Description: libs.Config.Admin.Rolename,
		Model:       gorm.Model{CreatedAt: time.Now()},
		Perms:       perms,
	}
	if err := easygorm.GetEasyGormDb().Create(&role).Error; err != nil {
		panic(fmt.Sprintf("seeder data create role err：%+v\n", err))
	}

	err := drole.AddPermForRole(role)
	if err != nil {
		panic(fmt.Sprintf("添加角色失败：%+v", err))
	}

	fmt.Println(fmt.Sprintf("填充角色：%+v", role.Name))
	fmt.Println(fmt.Sprintf("\n填充角色权限："))
	for _, perm := range role.Perms {
		fmt.Println(fmt.Sprintf("  %+v", perm))
	}
}

type Role struct {
	Id   uint
	Name string
}

// CreateAdmin 新建管理员
func CreateAdmin() {
	var roleIds []uint
	var roleNames []string
	var roles []*Role
	easygorm.GetEasyGormDb().Model(&models.Role{}).Find(&roles)
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
		roleNames = append(roleNames, role.Name)
	}

	//先创建user，客户client，管理员admin，的基本信息
	user := &models.User{
		RealName: "",
		Nickname: "",
		Phone: "18888888888",
		Birthday: time.Now().Unix(),
		CardId: "",
		Mark: "",
		Avatar: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132",
		Address: "厦门市湖里区金海湾1号楼B栋",
	}

	easygorm.GetEasyGormDb().Create(user)

	// 创建管理员的信息表
	admin := &madmin.Admin{
		MerchId: 1,
		Uid: 1,
		RealName: "宅职社",
		Nickname: "益帮手",
		Phone: "18888888888",
		Email: "xmxlb@163.com",
		Avatar: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132",
		UseApp: 1,
		UsePc: 1,
		Type: 1,
		Status: 1,
		LastIp: "",
		LastTime: time.Now().Unix(),
	}

	easygorm.GetEasyGormDb().Create(admin)

	account := &models.Account{
		Username: libs.Config.Admin.Username,
		Password: libs.HashPassword(libs.Config.Admin.Password),
		Identity: 1,  // 客户表client_id或管理员表admin_id
		Type: 1,        // 是否为管理员
		Status: 1,
		Phone: "18888888888",
		Email: "xmxlb@163.com",
		CreatedAt: time.Now().Unix(),
		CreatedUid: 1,
		UpdatedAt: time.Now().Unix(),
		UpdatedUid: 1,
		Effect: 0,
		Memo: "超级弱鸡程序猿一枚！！！！",

		RoleIds:  roleIds,
	}

	easygorm.GetEasyGormDb().Create(account)

	err := daccount.AddRoleForUser(account)
	if err != nil {
		panic(fmt.Sprintf("添加管理员失败：%+v", err))
	}

	fmt.Println(fmt.Sprintf("管理员密码：%s", libs.Config.Admin.Password))
	fmt.Println(fmt.Sprintf("管理员角色：%+v", roleNames))
}

//AutoMigrates 重置数据表
//easygorm.Egm.Db.DropTableIfExists 删除存在数据表
func AutoMigrates() {
	if err := DropTables(); err != nil {
		fmt.Println(fmt.Sprintf("seeder data  auto migrate  err：%+v\n", err))
		return
	}
	if err := easygorm.Migrate([]interface{}{
		&models.Role{},
		&models.Permission{},
		&models.Config{},

	}); err != nil {
		fmt.Println(fmt.Sprintf("seeder data  auto migrate  err：%+v\n", err))
		return
	}
}

// DropTables 删除数据表
func DropTables() error {
	prefix := libs.Config.DB.Prefix
	err := easygorm.GetEasyGormDb().Migrator().DropTable(
		prefix+"roles",
		prefix+"permissions",
		prefix+"configs",
	)

	if err != nil {
		return err
	}
	return nil
}
