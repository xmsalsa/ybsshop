/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-13 14:16:12
 */
package label

type SSaveSetLabels struct {
	ClientIds  []int `form:"client_ids" json:"client_ids" validate:"required,gt=0,unique,dive,gt=0" comment:"客户ID"`
	LabelIds   []int `form:"label_ids" json:"label_ids" validate:"required,gt=0,unique,dive,gte=0" comment:"标签ID"`
	MerchId    int   `form:"merch_id" json:"merch_id" validate:"omitempty,gte=0" comment:"商户id"`
	UpdatedUid int   `form:"updated_uid", json:"updated_uid"`
}

type SSetClientLabel struct {
	ClientId   int   `form:"client_id" json:"client_id" validate:"required,gt=0" comment:"客户ID"`
	LabelIds   []int `form:"label_ids" json:"label_ids" validate:"omitempty,gte=0,unique,dive,gt=0" comment:"标签ID"`
	MerchId    int   `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int   `form:"updated_uid", json:"updated_uid"`
}

type SClientLabelCount struct {
	ClientId int `form:"client_id" json:"client_id" validate:"required_without=ClientId,gt=0" comment:"客户ID"`
	LabelId  int `form:"label_ids" json:"label_ids" validate:"required_without=ClientId,gt=0" comment:"标签ID"`
	MerchId  int `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
}

type SClientLabelTree struct {
	ClientId int `form:"client_id" json:"client_id" validate:"required,gt=0" comment:"客户ID"`
	MerchId  int `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
}
