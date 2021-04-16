/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-15
 * Time: 下午 16:57
 */
package core

import (
	"encoding/json"
	"reflect"
	"sync"
)

type Event struct {
	Wg     sync.WaitGroup
	Name   string
	Params map[string]interface{}
	Data   map[string]interface{}
	IsErr  error
}

/**   暂不支持    **/
func (this *Event) GetParams(keys string, tem interface{}) {
	bVal := reflect.ValueOf(this.Params[keys]).Elem() //获取reflect.Type类型
	cVal := reflect.ValueOf(tem)
	if ok := bVal.FieldByName(keys).IsValid(); ok {
		bVal.FieldByName(keys).Set(reflect.ValueOf(this.Params[keys]).Convert(cVal.Type()))
	}
}

func (this *Event) GetData(keys string, tem interface{}) {
	strRet, err := json.Marshal(this.Data[keys])
	if err != nil {
	}
	// json转map
	err1 := json.Unmarshal(strRet, &tem)
	if err1 != nil {
	}
}

func (this *Event) AddData(keys string, data interface{}) {
	if len(this.Data) == 0 {
		this.Data = make(map[string]interface{})
	}
	this.Data[keys] = data
}

func (this *Event) AddParams(keys string, data interface{}) {
	if len(this.Params) == 0 {
		this.Params = make(map[string]interface{})
	}
	this.Params[keys] = data
}

func (ser *Event) Error() string {
	if ser.IsErr != nil {
		return ser.IsErr.Error()
	}
	return ""
}
