/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-15 01:13:39
 */
package account

type SAccountExsit struct {
	Username string `form:"username" json:"username" validate:"required,alphanumunicode,min=2,max=32" comment:"账号"`
}

type SAddAccount struct {
	Username   string `form:"username" json:"username" validate:"required,alphanumunicode,min=2,max=32" comment:"账号"`
	Password   string `form:"password" json:"password" validate:"omitempty,min=3" comment:"密码"`
	Identity   int    `form:"identity" json:"identity" validate:"required,gt=0" comment:"账号关联的ID"` // 客户表client_id或管理员表admin_id
	Phone      string `form:"phone" json:"phone"; validate:"omitempty" comment('手机号码');`
	Email      string `form:"email" json:"email"; validate:"omitempty,max=128" comment('邮箱');');`
	Type       int    `form:"type" json:"type" validate:"omitempty,gte=0" comment('账号类型:1管理员;2客户');" `
	Status     int    ` form:"status" json:"status" validate:"omitempty,gte=0" comment('1正常，禁止,2注销')" json:"status"`
	UpdatedUid int    `json:"updated_uid"`
}
