/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-15
 * Time: 下午 21:36
 */
package product

type StoreProductReply struct {
	Id                   int     `gorm:"not null pk autoincr comment('评论ID') INT(11)"`
	Uid                  int     `gorm:"not null comment('用户ID') index INT(11)"`
	MerchId              int     `gorm:"not null comment('商户 id') INT(10)"`
	OrderProductId       int     `gorm:"not null comment('订单商品ID') unique(order_id) INT(11)"`
	Unique               string  `gorm:"comment('商品 sku') unique(order_id) CHAR(12)"`
	ProductId            int     `gorm:"not null comment('商品id') index INT(11)"`
	ProductType          int     `gorm:"not null default 0 comment('0=普通商品') TINYINT(4)"`
	ProductScore         int     `gorm:"not null comment('商品分数') TINYINT(1)"`
	ServiceScore         int     `gorm:"not null comment('服务分数') TINYINT(1)"`
	PostageScore         int     `gorm:"not null comment('物流分数') TINYINT(1)"`
	Rate                 float32 `gorm:"default 5.0 comment('平均值') FLOAT(2,1)"`
	Comment              string  `gorm:"not null comment('评论内容') VARCHAR(512)"`
	Pics                 string  `gorm:"not null comment('评论图片') TEXT"`
	MerchantReplyContent string  `gorm:"comment('管理员回复内容') VARCHAR(300)"`
	MerchantReplyTime    int64   `gorm:"comment('管理员回复时间') INT"`
	IsDel                int     `gorm:"not null default 0 comment('0未删除1已删除') TINYINT(1)"`
	IsReply              int     `gorm:"not null default 0 comment('0未回复1已回复') TINYINT(1)"`
	IsVirtual            int     `gorm:"not null default 0 comment('0不是虚拟评价1是虚拟评价') TINYINT(1)"`
	Nickname             string  `gorm:"not null comment('用户名称') VARCHAR(64)"`
	Avatar               string  `gorm:"not null comment('用户头像') VARCHAR(255)"`
	CreatedAt            int64   `gorm:"not null default 0 INT(11)" json:"created_at"`
	CreatedUid           int64   `gorm:"not null default 0 INT(11)" json:"created_user_id"`
	UpdatedAt            int64   `gorm:"not null default 0 INT(11)" json:"updated_at"`
	UpdatedUid           int64   `gorm:"not null default 0 INT(11)" json:"updated_user_id"`
	Effect               int64   `gorm:"not null default 1  TINYINT(1)" json:"effect"`
}
