/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-07 10:23:43
 */
package libs

import (
	"encoding/json"
	"reflect"
	"strings"
)

/**
 * 结构体转成json 字符串
 * @method StruckToString
 * @param  {[type]}       data interface{} [description]
 */
func StructToString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}

/**
 * 结构体转换成map对象
 * @method func
 * @param  {[type]} t *Tools        [description]
 * @return {[type]}   [description]
 */
func StructToMap(obj interface{}) map[string]interface{} {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < k.NumField(); i++ {
		data[strings.ToLower(k.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

// 结构体拷贝
func StructCopy(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)            // 值
		name := sval.Type().Field(i).Name // 键

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		// 按键值类型转化, 不同类型时, 值可能产生变化  未确定
		dvalue.Set(reflect.ValueOf(value.Interface()))
	}
}
