/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-07 10:23:43
 */
package utils

import (
	"reflect"
)

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
