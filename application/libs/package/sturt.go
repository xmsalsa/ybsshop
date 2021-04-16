/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 09:17:34
 */
/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-8
 * Time: 下午 16:55
 */
package _package

import (
	"reflect"
	"strings"
)

//将结构体里的成员按照json名字来赋值
func SetStructFieldByJsonName(ptr interface{}, fields map[string]interface{}) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")

		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		name = strings.Split(name, ",")[0]
		if value, ok := fields[name]; ok {
			//给结构体赋值
			//保证赋值时数据类型一致
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}
	return
}

//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		s, _ := bVal.Type().FieldByName(name)
		if vVal.Type().Field(i).Type == s.Type {
			if ok := bVal.FieldByName(name).IsValid(); ok {
				bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
			}
		}
	}
}
