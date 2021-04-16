/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-15
 * Time: 下午 21:49
 */
package product

import (
	_package "shop/application/libs/package"
)

/* 定义结构体 */
type ProductReplysService struct {
	/* 错误体 */
	isErr error
}

/**	添加虚拟（回复/评论）	*/
func (ser *ProductReplysService) SaveFictitiousReply(param PostProduct, Authority _package.Authority) {
	//初始属性值

}
