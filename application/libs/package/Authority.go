package _package

import (
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
)

type Authority struct {
	MerchId int
	UserId  int
}

func GetAuthority(Ctx iris.Context) Authority {
	Authority := Authority{
		MerchId: 1,
		UserId:  1,
	}

	return Authority
}

func Sql(Ctx iris.Context) map[string]interface{} {
	Authority := GetAuthority(Ctx)
	merch := make([]interface{}, 0)
	auth := make(map[string]interface{})
	auth["merch_id"] = append(merch, "merch_id", "=", Authority.MerchId)
	return auth
}

func GetUpdateInit(Ctx iris.Context) map[string]interface{} {
	init := make(map[string]interface{})
	Authority := GetAuthority(Ctx)
	init["updated_at"] = time.Now().Unix()
	init["updated_uid"] = Authority.UserId
	return init
}

func AddStructCommon(id int, binding interface{}, Authority Authority) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	if ok := bVal.FieldByName("MerchId").IsValid(); ok {
		bVal.FieldByName("MerchId").Set(reflect.ValueOf(Authority.MerchId))
	}
	if ok := bVal.FieldByName("UpdatedAt").IsValid(); ok {
		bVal.FieldByName("UpdatedAt").Set(reflect.ValueOf(time.Now().Unix()))
	}
	if ok := bVal.FieldByName("UpdatedUid").IsValid(); ok {
		bVal.FieldByName("UpdatedUid").Set(reflect.ValueOf(Authority.UserId).Convert(bVal.FieldByName("UpdatedUid").Type()))
	}
	if id == 0 {
		if ok := bVal.FieldByName("CreatedAt").IsValid(); ok {
			bVal.FieldByName("CreatedAt").Set(reflect.ValueOf(time.Now().Unix()))
		}
		if ok := bVal.FieldByName("Effect").IsValid(); ok {
			bVal.FieldByName("Effect").Set(reflect.ValueOf(1).Convert(bVal.FieldByName("Effect").Type()))
		}
		if ok := bVal.FieldByName("CreatedUid").IsValid(); ok {
			bVal.FieldByName("CreatedUid").Set(reflect.ValueOf(Authority.UserId).Convert(bVal.FieldByName("CreatedUid").Type()))
		}
	}
}
