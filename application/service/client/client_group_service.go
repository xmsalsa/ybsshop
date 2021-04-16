/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-12 22:58:42
 */
package client

import (
	"errors"
	"fmt"
	"shop/application/libs/response"
	"shop/application/libs/utils"
	mclient "shop/application/models/client"
	mgroup "shop/application/models/client/group"
	"shop/application/service/client/group"
)

/*
 一些与用户相关的操作
 @module GROUP
*/

// 设置用户分组
func SaveSetGroup(param SSaveSetGroup) error {
	var err error
	if param.GroupID > 0 {
		gparam := group.SGroupDetail{
			Id:      param.GroupID,
			MerchId: param.MerchId,
		}
		var groups mgroup.ClientGroup
		groups, err := group.GroupDetail(gparam)
		if err != nil {
			return err
		}
		if groups.Id == 0 {
			return errors.New(fmt.Sprintf("%s%s", groups.TableComment(), response.DB_RECORD_NOEXSIT))
		}
	}

	condition := make(map[string]interface{})
	if param.MerchId > 0 {
		condition["merch_id"] = param.MerchId
	}
	data := utils.InitUpdate(param.UpdatedUid)
	data["group_id"] = param.GroupID
	db := utils.GetGormDbWithModel(mclient.Client{})
	if len(condition) > 0 {
		utils.Build(db, condition)
	}
	if len(param.ClientIds) == 1 {
		db.Where("id", param.ClientIds[0])
	} else {
		db.Where(param.ClientIds)
	}
	err = db.Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
