package utils

import (
	"time"
)

//GetCurrentDateTime 获取当前的日期时间
func GetCurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//GetCurrentUnixTime 获取当前的时间戳
func GetCurrentUnixTime() int64 {
	return time.Now().Unix()
}

//GetCurrentUnixNano 获取当前的纳秒时间戳
func GetCurrentUnixNano() int64 {
	return time.Now().Unix()
}

//GetDateTimeToUnix 日期时间转时间戳
func GetDateTimeToUnix(dateTime string) int64 {
	// 时区
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", dateTime, Loc)
	return t2.Unix()
}

//GetUnixTimeToDate 时间戳转日期时间
func GetUnixTimeToDate(timestamp int64) string {
	// 时区
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Unix(timestamp, 0).In(Loc).Format("2006-01-02 15:04:05")
}

func GetDaysByTwoDate(currentDate, beforeDate string) int64 {
	//指定日期转为时间戳
	current := GetDateTimeToUnix(currentDate)
	before := GetDateTimeToUnix(beforeDate)
	return int64((current - before) / 86400)
}
