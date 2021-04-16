/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 16:14:37
 */
package admin

type SAdminCreate struct {
	MerchId  int    `form:"merch_id" json:"merch_id" validate:"omitempty,gte=0" comment:"商户id"`
	Username string `form:"username" json:"username" validate:"required,alphanumunicode,min=2,max=32" comment:"账号"`
	Password string `form:"password" json:"password" validate:"required,alphanumunicode,min=3" comment:"密码"`

	RealName string `form:"real_name" json:"real_name"; validate:"omitempty,max=32" comment('真实姓名');`
	Nickname string `form:"nickname" json:"nickname"; validate:"omitempty,max=32" comment('昵称');`
	Email    string `form:"email" json:"email"; validate:"omitempty,max=128" comment('邮箱');');`
	Avatar   string `form:"avatar" json:"avatar" validate:"omitempty,max=255" comment('头像');'`

	Phone      string `form:"phone" json:"phone";comment('手机号码');`
	UseApp     int    `form:"use_app" json:"use_app" validate:"omitempty,gte=0" comment('使用APP:1是;0否');`
	UsePc      int    `form:"use_pc" json:"use_pc" validate:"omitempty,gte=0" comment('使用PC:1是;0否');" `
	Type       int    `form:"type" json:"type" validate:"omitempty,gte=0" comment('账号类型:1系统管理员;2商户管理员');" `
	Status     int    ` form:"status" json:"status" validate:"omitempty,gte=0" comment('1正常，0离职,')" json:"status"`
	UpdatedUid int    `json:"updated_uid"`
	// 权限角色 略
}

type SAddAdmin struct {
	Username string `form:"username" json:"username" validate:"required,alphanumunicode,min=2,max=32" comment:"账号"`
	Password string `form:"password" json:"password" validate:"omitempty,min=3" comment:"密码"`
	MerchId  int    `form:"merch_id" json:"merch_id" validate:"omitempty,gte=0" comment:"商户id"`
	Uid      int    `form:"uid" json:"uid" validate:"omitempty,gte=0" comment:"用户id"`

	RealName string `form:"real_name" json:"real_name"; validate:"omitempty,max=32" comment('真实姓名');`
	Nickname string `form:"nickname" json:"nickname"; validate:"omitempty,max=32" comment('昵称');`
	Email    string `form:"email" json:"email"; validate:"omitempty,max=128" comment('邮箱');');`
	Avatar   string `form:"avatar" json:"avatar" validate:"omitempty,max=255" comment('头像');'`

	Phone      string `form:"phone" json:"phone"; validate:"omitempty" comment('手机号码');`
	UseApp     int    `form:"use_app" json:"use_app" validate:"omitempty,gte=0" comment('使用APP:1是;0否');`
	UsePc      int    `form:"use_pc" json:"use_pc" validate:"omitempty,gte=0" comment('使用PC:1是;0否');" `
	Type       int    `form:"type" json:"type" validate:"omitempty,gte=0" comment('账号类型:1系统管理员;2商户管理员');" `
	Status     int    ` form:"status" json:"status" validate:"omitempty,gte=0" comment('1正常,0离职,')" json:"status"`
	UpdatedUid int    `json:"updated_uid"`
}
