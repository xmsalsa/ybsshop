/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 11:27:34
 */
package client

type SSaveSetGroup struct {
	ClientIds  []int `form:"client_ids" json:"client_ids" validate:"required,gt=0,dive,gt=0" comment:"客户ID"`
	GroupID    int   `form:"group_id" json:"group_id" validate:"omitempty,gt=0" comment:"分组ID"`
	MerchId    int   `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int   `form:"updated_uid", json:"updated_uid"`
}
