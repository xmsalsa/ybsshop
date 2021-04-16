package label

/************* 标签 *************/
// 标签创建  validate:"required,numeric,gte=1"
type SLabelCreate struct {
	Id          int    `form:"id" json:"id" comment:"标签id"`
	MerchId     int    `form:"merch_id" json:"merch_id" comment:"商户id"`
	Name        string `form:"label_name" json:"label_name" validate:"required,min=1" comment:"标签名字"`
	LabelcateId int    `form:"labelcate_id" json:"labelcate_id" validate:"required,gte=1" comment:"标签分类ID"`
	Sort        int    `form:"sort" json:"sort" comment:"标签排序"`
	Type        int    `form:"type" json:"type" comment:"1用户标签"`
	Source      int    `form:"source" json:"source" comment:"0=手动标签 1=自动标签"`
	UpdatedUid  int    `json:"updated_uid"`
}

type SLabelDetail struct {
	Id         int    `form:"id" json:"id" validate:"required,numeric,gte=1" comment:"标签id"`
	MerchId    int    `form:"merch_id" json:"merch_id" comment:"商户id"`
	Name       string `form:"label_name" json:"label_name" comment:"标签名字"`
	UpdatedUid int    `json:"updated_uid"`
}
type SLabelModify struct {
	Id          int    `form:"id" json:"id" validate:"required,numeric,gte=1" comment:"标签id"`
	MerchId     int    `form:"merch_id" json:"merch_id" comment:"商户id"`
	Name        string `form:"label_name" json:"label_name" validate:"required,min=1" comment:"标签名字"`
	LabelcateId int    `form:"labelcate_id" json:"labelcate_id" validate:"gte=1" comment:"标签分类ID"`
	Sort        int    `form:"sort" json:"sort" comment:"标签排序"`
	Type        int    `form:"type" json:"type" comment:"1用户标签"`
	Source      int    `form:"source" json:"source" comment:"0=手动标签 1=自动标签"`
	UpdatedUid  int    `form:"updated_uid" json:"updated_uid" validate:"" comment:""`
}

// 标签分页参数
type SLabelAll struct {
	MerchId      int   `form:"merch_id" json:"merch_id" validate:"min=0" comment:"商户id"`
	LabelcateId  int   `form:"labelcate_id" json:"labelcate_id" validate:"min=0" comment:"标签分类id"`
	LabelcateIds []int `form:"labelcate_ids" json:"labelcate_ids" validate:"omitempty,unique,dive,gt=0" comment:"标签分类id"`
}

/************* 标签分类 *************/
type SLabelcateCreate struct {
	Id         int    `form:"id" json:"id" comment:"标签id"`
	Name       string `form:"labelcate_name" json:"labelcate_name" validate:"required,min=1" comment:"标签分类名字"`
	Sort       int    `form:"sort" json:"sort" validate:"required,numeric,gte=1" comment:"标签分类排序"`
	Type       int    `form:"type" json:"type" validate:"required,numeric,gte=1" comment:"标签分类排序"`
	MerchId    int    `form:"merch_id" json:"merch_id" comment:"商户id"`
	UpdatedUid int    `json:"updated_uid"`
}

// 标签分类详情参数
type SLabelcateDetail struct {
	Id         int `form:"id" json:"id" validate:"numeric,min=0,required" comment:"标签分类id"`
	MerchId    int `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	UpdatedUid int `json:"updated_uid"`
}

type SLabelcateUpdate struct {
	Id         int    `form:"id" json:"id" validate:"numeric,min=0,required" comment:"标签分类id"`
	MerchId    int    `form:"merch_id" json:"merch_id" validate:"gte=0" comment:"商户id"`
	Name       string `form:"labelcate_name" json:"labelcate_name" validate:"required,min=1" comment:"标签分类名字"`
	Sort       int    `form:"sort" json:"sort" validate:"required,numeric,gte=1" comment:"标签分类排序"`
	UpdatedUid int    `json:"updated_uid"`
}

// 标签分类分页参数
type SLabelcateAll struct {
	Ids     []int `form:"labelcate_ids" json:"labelcate_ids" validate:"omitempty,unique,dive,gt=0" comment:"标签分类ID"`
	MerchId int   `form:"merch_id" json:"merch_id" validate:"min=0" comment:"商户id"`
}
