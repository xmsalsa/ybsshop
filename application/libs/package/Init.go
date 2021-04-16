package _package

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const INITCODE = 200
const LIMIT = 15
const PAGE = 1
const EFFECT = 1
const EFFECTD = 0
const ACTIVE = 1
const ACTIVED = 0

func EmptyDta() map[string]interface{} {
	return make(map[string]interface{})
}

func DeleteKeys(effect int, active int) string {
	h := md5.New()
	h.Write([]byte(IntToBytes(effect + active)))
	return hex.EncodeToString(h.Sum(nil))
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

// JSON序列化方式
func JsonStructToMap(stuObj interface{}) (map[string]interface{}, error) {
	// 结构体转json
	strRet, err := json.Marshal(stuObj)
	if err != nil {
		return nil, err
	}
	// json转map
	var mRet map[string]interface{}
	err1 := json.Unmarshal(strRet, &mRet)
	if err1 != nil {
		return nil, err1
	}
	return mRet, nil
}

//时间戳转时间
func UnixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

/******************** 时间日期 START *********************/
const (
	DATETIME_FORMAT       = "2006-01-02 03:04:05" // 日期时间
	DATETIME_FORMAT_SHORT = "2006-01-02"          // 日期
	TIME_ZONE             = "Asia/Shanghai"       // 时区
)

/**
 * 格式化时间
 */
func FormatDatetime(time time.Time) string {
	return time.Format(DATETIME_FORMAT)
}

//时间戳转时间
func UnixToDatetime(timeUnix int64) string {
	timeStr := time.Unix(timeUnix, 0).Format(DATETIME_FORMAT)
	return timeStr
}

//时间戳转日期
func UnixToDate(timeUnix int64) string {
	timeStr := time.Unix(timeUnix, 0).Format(DATETIME_FORMAT_SHORT)
	return timeStr
}

//时间转时间戳
func DatetimeToUnix(timeStr string) (int64, error) {
	local, err := time.LoadLocation(TIME_ZONE) //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(TIME_ZONE, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}

/******************** 时间日期  END  *********************/

func F2i(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}

func F2si(f string) int {
	Price, _ := strconv.ParseFloat(f, 64)
	return F2i(Price * 100)
}

func IF64Fl64(data interface{}) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", data), 64)
	return value
}

func IF64Str(data int, digit int) string {
	return fmt.Sprintf("%.2f", float64(data)/float64(digit))
}

func Int64Fint(nums int64) int {
	f := strconv.FormatInt(nums, 10)
	int, _ := strconv.Atoi(f)
	return int
}

func GetMapKey(key string, v []map[string]interface{}) string {
	str := ""
	for _, s := range v {
		str = fmt.Sprintf("%s%s,", str, s[key])
	}
	return strings.Trim(str, ",")
}

func SlicesStrFInt(data []string) []int {
	f := make([]int, 0)
	for _, s := range data {
		i, _ := strconv.Atoi(s)
		f = append(f, i)
	}
	return f
}
