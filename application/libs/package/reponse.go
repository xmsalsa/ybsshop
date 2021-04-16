package _package

import (
	"github.com/kataras/iris/v12/mvc"
	"reflect"
)

func List(data interface{}, count int64) map[string]interface{} {
	if InterfaceIsNil1(data) == true && InterfaceIsNil2(data) == true {
		data = make([]interface{}, 0)
	}
	return map[string]interface{}{"list": data, "count": count}
}

func Response(code int64, msg string, data interface{}) mvc.Result {
	return mvc.Response{
		Object: Test(code, msg, data),
	}
}

func Test(code int64, msg string, data interface{}) map[string]interface{} {
	if InterfaceIsNil1(data) == true && InterfaceIsNil2(data) == true {
		data = [][]int64{}
	}
	return map[string]interface{}{
		"msg":    msg,
		"data":   data,
		"status": code,
	}
}

//异常判断
func InterfaceIsNil1(i interface{}) bool {
	ret := i == nil

	if !ret { //需要进一步做判断
		defer func() {
			recover()
		}()
		ret = reflect.ValueOf(i).IsNil() //值类型做异常判断，会panic的
	}

	return ret
}

//类型判断
func InterfaceIsNil2(i interface{}) bool {
	ret := i == nil

	if !ret { //需要进一步做判断
		vi := reflect.ValueOf(i)
		kind := reflect.ValueOf(i).Kind()
		if kind == reflect.Slice ||
			kind == reflect.Map ||
			kind == reflect.Chan ||
			kind == reflect.Interface ||
			kind == reflect.Func ||
			kind == reflect.Ptr {
			return vi.IsNil()
		}
	}

	return ret
}
