/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-12 08:48:37
 */
package group

/************* 用户分组 *************/
type SGroupCreate struct {
	Id         int    `form:"id" json:"id" comment:"id"`
	Name       string `form:"group_name" json:"group_name" validate:"required,min=1" comment:"用户分组名字"`
	Sort       int    `form:"sort" json:"sort" validate:"required,numeric,gte=1" comment:"用户分组排序"`
	MerchId    int    `form:"merch_id" json:"merch_id" comment:"商户id"`
	UpdatedUid int    `json:"updated_uid"`
}

type SGroupDetail struct {
	Id         int `form:"id" json:"id" validate:"omitempty,gte=0" comment:"id"`
	MerchId    int `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int `json:"updated_uid"`
}

type SGroupUpdate struct {
	Id         int    `form:"id" json:"id" validate:"numeric,min=0,required" comment:"id"`
	MerchId    int    `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	Name       string `form:"group_name" json:"group_name" validate:"required,min=1" comment:"用户分组名字"`
	Sort       int    `form:"sort" json:"sort" validate:"numeric,gte=1" comment:"用户分组排序"`
	UpdatedUid int    `json:"updated_uid"`
}

type SGroupAll struct {
	MerchId int `form:"merch_id" json:"merch_id" validate:"min=0" comment:"商户id"`
}

type SSetGroup struct {
	GroupId    int   `form:"group_id" json:"group_id" validate:"numeric,min=0,required" comment:"用户分组ID"`
	ClientIds  []int `form:"client_ids" json:"client_ids" validate:"min=0,required" comment:"用户ID"`
	MerchId    int   `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int   `json:"updated_uid"`
}
