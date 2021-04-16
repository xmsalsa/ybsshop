package utils

import "time"

/**
 * 时间格式化
 * @method func
 * @param  {[type]} t *Tools        [description]
 * @return {[type]}   [description]
 */
func TimeFormat(time *time.Time, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Format(format)
}

const (
	DATETIME_FORMAT       = "2006-01-02 15:04:05" // 日期时间
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
