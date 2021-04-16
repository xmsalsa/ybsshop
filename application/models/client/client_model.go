/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-12 22:17:31
 */
package client

import (
	"shop/application/libs/utils"
	"shop/application/models"
)

type Client struct {
	models.Model
	Uid      int    `gorm:"type:INT; default:0; column:uid; comment('用户ID')" json:"uid"`
	MerchId  int    `gorm:"type:INT; default:0; column:merch_id; comment('商户ID')" json:"merch_id"`
	RealName string `gorm:"type:VARCHAR(24); default:'';column:real_name;comment('真实姓名'); " json:"real_name"`
	NickName string `gorm:"type:VARCHAR(64); default:'';column:nick_name;comment('昵称'); " json:"nick_name"`
	Phone    string `gorm:"type:VARCHAR(15); default:'';column:phone;comment('手机号码'); " json:"phone"`
	Email    string `gorm:"type:VARCHAR(128); default:'';column:email;comment('邮箱'); " json:"email"`
	Birthday int64  `gorm:"type:INT; default:0; column:birthday; comment('生日')" json:"birthday"`
	CardId   string `gorm:"type:VARCHAR(24); default:'';column:card_id;comment('身份证号码'); " json:"card_id"`
	Mark     string `gorm:"type:VARCHAR(255); default:'';column:mark;comment('客户备注'); " json:"mark"`
	Avatar   string `gorm:"type:VARCHAR(255); default:'';column:avatar;comment('头像'); " json:"avatar"`
	Address  string `gorm:"type:VARCHAR(255); default:'';column:address;comment('详细地址'); " json:"address"`

	PartnerId   int   `gorm:"type:INT; default:0; column:partner_id; comment('合伙人ID')" json:"partner_id"`
	GroupId     int   `gorm:"type:INT; default:0; column:group_id; comment('分组ID')" json:"group_id"`
	SpreadUid   int   `gorm:"type:INT; default:0; column:spread_uid; comment('推广员ID')" json:"spread_uid"`
	SpreadTime  int64 `gorm:"type:INT; default:0; column:spread_time; comment('推广员关联时间')" json:"spread_time"`
	SpreadCount int   `gorm:"type:INT; default:0; column:spread_count; comment('下级数量')" json:"spread_count"`
	IsPromoter  int   `gorm:"type:INT; default:0; column:is_promoter; comment('是否为推广员:1是;0否')" json:"is_promoter"`

	Source       int    `gorm:"not null; type:tinyint(1); default:1; column:source; comment('来源1pc,2app,3h5,4minip')" json:"source"`
	Type         string `gorm:"type:VARCHAR(32); default:1; column:type; comment('用户类型')" json:"type"`
	Status       int    `gorm:"type:INT; default:1; column:status; comment('1为正常，0为禁止')" json:"status"`
	Level        int    `gorm:"type:INT; default:0; column:level; comment('等级')" json:"level"`
	CleanTime    int64  `gorm:"type:INT; default:0; column:clean_time; comment('清理会员时间')" json:"clean_time"`
	IsMoneyLevel int    `gorm:"type:INT; default:0; column:is_money_level; comment('会员来源 0: 购买商品升级 1：花钱购买的会员2: 会员卡领取')" json:"is_money_level"`
	isEverLevel  int    `gorm:"type:INT; default:0; column:is_ever_level; comment('是否永久性会员 0: 非永久会员 1：永久会员')" json:"is_ever_level"`
	OverdueTime  int64  `gorm:"type:INT; default:0; column:is_ever_level; comment('会员到期时间')" json:"is_ever_level"`

	Balance   int `gorm:"type:INT; default:0; column:balance; comment('余额')" json:"balance"`
	Brokerage int `gorm:"type:INT; default:0; column:brokerage; comment('佣金金额')" json:"brokerage"`
	Integral  int `gorm:"type:INT; default:0; column:integral; comment('用户剩余积分')" json:"integral"`
	Exp       int `gorm:"type:INT; default:0; column:exp; comment('会员经验')" json:"exp"`
	PayCount  int `gorm:"type:INT; default:0; column:pay_count; comment('支付次数')" json:"pay_count"`

	AddTime     int64  `gorm:"type:INT; default:0; column:add_time; comment('添加时间')" json:"add_time"`
	AddIp       string `gorm:"type:VARCHAR(16); default:'';column:add_ip;comment('添加IP'); " json:"add_ip"`
	LastTime    int64  `gorm:"type:INT; default:0; column:last_time; comment('最后一次登录时间')" json:"last_time"`
	LastIp      string `gorm:"type:VARCHAR(16); default:'';column:last_ip;comment('最后一次登录ip'); " json:"last_ip"`
	LoginType   string `gorm:"type:VARCHAR(36); default:'';column:login_type;comment('用户登陆类型，h5,wechat,routine'); " json:"login_type"`
	RecordPhone string `gorm:"type:VARCHAR(36); default:'';column:record_phone;comment('记录临时电话'); " json:"record_phone"`
	SignNum     int    `gorm:"type:INT; default:0; column:sign_num; comment('连续签到天数')" json:"sign_num"`
}

func (this *Client) TableComment() string {
	return "客户"
}

func (this *Client) RepsToDesc() map[string]interface{} {
	if this.Id == 0 {
		return map[string]interface{}{"id": 0}
	} else {
		return map[string]interface{}{
			"id":         this.Id,
			"merch_id":   this.MerchId,
			"created_at": utils.UnixToDatetime(this.CreatedAt),
		}
	}
}

// 创建用户时, 初始基本信息
func (this *Client) Init() Client {
	return Client{
		PartnerId:   0, //合伙人ID
		GroupId:     0, //分组ID
		SpreadUid:   0, //推广员ID
		SpreadTime:  0, //推广员关联时间
		SpreadCount: 0, //下级数量
		IsPromoter:  0, //是否为推广员:1是;0否

		Source:       1,   //来源1pc,2app,3h5,4minip
		Type:         "1", // 用户类型
		Status:       1,   // 1为正常，0为禁止
		Level:        0,   //等级
		CleanTime:    0,   //清理会员时间
		IsMoneyLevel: 0,   //会员来源 0: 购买商品升级 1：花钱购买的会员2: 会员卡领取
		isEverLevel:  0,   //是否永久性会员 0: 非永久会员 1：永久会员
		OverdueTime:  0,   //会员到期时间

		Balance:   0, //余额
		Brokerage: 0, //佣金金额
		Integral:  0, //用户剩余积分
		Exp:       0, //会员经验
		PayCount:  0, //支付次数

		AddTime:     0,  //添加时间
		AddIp:       "", //添加IP
		LastTime:    0,  //最后一次登录时间
		LastIp:      "", //最后一次登录ip
		LoginType:   "", //用户登陆类型，h5,wechat,routine
		RecordPhone: "", //记录临时电话
		SignNum:     0,  //连续签到天数

	}
}
