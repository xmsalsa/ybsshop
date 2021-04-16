package service

/************* 参数结构 *************/
// 创建
type SCreate struct {
	Id         int64 `form:"id" json:"id" comment:"id"`
	MerchId    int64 `form:"merch_id" json:"merch_id" comment:"商户id"`
	UpdatedUid int64 `json:"updated_uid"`
}

// 详情
type SDetail struct {
	Id         int64 `form:"id" json:"id" validate:"numeric,min=0,required" comment:"id"`
	MerchId    int64 `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int64 `json:"updated_uid"`
}

// 更新
type SUpdate struct {
	Id         int64 `form:"id" json:"id" validate:"numeric,min=0,required" comment:"id"`
	MerchId    int64 `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int64 `json:"updated_uid"`
}

// 分页
type SAll struct {
	MerchId int64 `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
}
