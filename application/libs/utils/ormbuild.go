/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-11 13:16:53
 */
package utils

import (
	"reflect"
	"shop/application/libs/easygorm"
	"time"

	"gorm.io/gorm"
)

// id 数据库统一是 int64, 时间为时间戳 int64
func InitModel(obj interface{}, uid, merch_id int) {
	udpated_at := time.Now().Unix()

	bVal := reflect.ValueOf(obj).Elem() //获取reflect.Type类型
	if ok := bVal.FieldByName("MerchId").IsValid(); ok {
		bVal.FieldByName("MerchId").Set(reflect.ValueOf(merch_id))
	}
	if ok := bVal.FieldByName("CreatedAt").IsValid(); ok {
		bVal.FieldByName("CreatedAt").Set(reflect.ValueOf(udpated_at))
	}
	if ok := bVal.FieldByName("CreatedUid").IsValid(); ok {
		bVal.FieldByName("CreatedUid").Set(reflect.ValueOf(uid))
	}
	if ok := bVal.FieldByName("UpdatedAt").IsValid(); ok {
		bVal.FieldByName("UpdatedAt").Set(reflect.ValueOf(udpated_at))
	}
	if ok := bVal.FieldByName("UpdatedUid").IsValid(); ok {
		bVal.FieldByName("UpdatedUid").Set(reflect.ValueOf(uid))
	}
	if ok := bVal.FieldByName("Effect").IsValid(); ok {
		bVal.FieldByName("Effect").Set(reflect.ValueOf(int(1)))
	}
}

// 更新时基础数据
func InitUpdate(uid int) map[string]interface{} {
	return map[string]interface{}{
		"updated_at":  time.Now().Unix(),
		"updated_uid": uid,
	}
}

// 使用链式操作
func GetGormDbWithModel(model interface{}) *gorm.DB {
	return easygorm.GetEasyGormDb().Model(model)
}

func GetGormDb() *gorm.DB {
	return easygorm.GetEasyGormDb()
}

// 丰富一点呀
func Build(db *gorm.DB, condition map[string]interface{}) {
	// 基础操作
	if _, ok := condition["id"]; ok {
		if IsSlice(condition["id"]) {
			db.Where(condition["id"])
		} else {
			db.Where("id=?", condition["id"])
		}
	}
	if _, ok := condition["merch_id"]; ok {
		db.Where("merch_id=?", condition["merch_id"])
	}
	var effect interface{}
	if _, ok := condition["effect"]; ok {
		effect = condition["effect"]
	} else {
		effect = 1
	}
	db.Where("effect=?", effect)

	// 这一步要搞个高级的, 注:condition 无序...尴尬
	for k, v := range condition {
		switch k {
		case "or":
			or(db, v)

		case "id":
			// 啥也不做
		case "merch_id":
			// 啥也不做
		case "effect":
		// 啥也不做
		default:
			if IsSlice(v) {
				db.Where(k+" in (?)", v)
			} else {
				db.Where(k, v)
			}
		}
	}
	return
}

// build or
func or(db *gorm.DB, condition interface{}) {

	return
}

// build not
func not(db *gorm.DB, condition interface{}) {
	return
}

// 分页情况
func BuildPage(db *gorm.DB, page Pages) {
	if page.Offset > 0 {
		db.Offset(int(page.Offset))
	}
	if page.Limit > 0 {
		db.Limit(int(page.Limit))
	}

	return
}
