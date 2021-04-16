package utils

import (
	"encoding/json"
	"strconv"
)

/*
 数字类型相关
*/
func ToInt64(i interface{}) int64 {
	//fmt.Printf(" ToInt64 %v 原类型: %T \n", i, i)
	switch t := i.(type) {
	default:
		return 0
	case nil:
		return 0
	case int:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case uint:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case json.Number:
		num, err := t.Int64()
		if err != nil {
			return 0
		}
		return int64(num)
	case string:
		num, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return 0
		}
		return num
	}
}

func ToString(i interface{}) string {
	var key string
	if i == nil {
		return key
	}

	switch i.(type) {
	case float64:
		ft := i.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := i.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := i.(int)
		key = strconv.Itoa(it)
	case uint:
		it := i.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := i.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := i.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := i.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := i.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := i.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := i.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := i.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := i.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = i.(string)
	case []byte:
		key = string(i.([]byte))
	default:
		newValue, _ := json.Marshal(i)
		key = string(newValue)
	}

	return key
}

// 是 slice 返回true, 否false
func IsSlice(i interface{}) bool {
	var is bool
	switch i.(type) {
	case []int:
		is = true
	case []int64:
		is = true
	case []string:
		is = true
	default:
		is = false
	}

	return is
}

//InArrayInt64 如果 i在 items 中,返回 true；否则，返回 false。
func InArrayInt64(items []int64, i int64) bool {
	for _, item := range items {
		if item == i {
			return true
		}
	}
	return false
}

func InArrayInt(items []int, i int) bool {
	for _, item := range items {
		if item == i {
			return true
		}
	}
	return false
}
