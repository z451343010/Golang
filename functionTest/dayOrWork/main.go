package main

import (
	"fmt"
	"time"
)

func dayOrWork(day int64) {
	var remain int64 = day % 5
	switch remain {
	case 0, 1, 2:
		fmt.Println("打鱼")
	default:
		fmt.Println("晒网")
	}
}

// 把日期字符串转换成unix时间戳
func timeStrToUnixStamp(timeStr string) (timestamp int64) {

	// 获取时区
	loc, _ := time.LoadLocation("Local")
	// Go系统内置的初始化时间模板
	timeLayout := "2006-01-02 15:04:05"
	// 转换成日期
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc)
	// 转换成时间戳
	timestamp = theTime.Unix()

	return timestamp

}

func main() {

	startTimeStr := "1990-01-01 00:00:00"
	endTimeStr := "1991-01-04 00:00:00"

	var timeSub int64
	timeSub = int64((timeStrToUnixStamp(endTimeStr) -
		timeStrToUnixStamp(startTimeStr)) / 24 / 60 / 60)

	dayOrWork(timeSub)

}
