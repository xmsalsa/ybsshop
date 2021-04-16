/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-16
 * Time: 上午 9:17
 */
package product

type StoreProductDescription struct {
	ProductId   int    `gorm:"not null default 0 comment('商品ID') index(product_id) INT(11)"`
	Description string `gorm:"not null comment('商品详情') TEXT"`
	Type        int    `gorm:"not null default 0 comment('商品类型') index(product_id) TINYINT(1)"`
}
