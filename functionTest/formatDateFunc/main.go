package main

import (
	"fmt"
	"time"
)

// 格式化日期和时间
func main() {

	// 方式一
	timeNow := time.Now()
	fmt.Printf("当前年月日:%02d-%02d-%02d %02d:%02d:%02d\n",
		timeNow.Year(), timeNow.Month(), timeNow.Day(),
		timeNow.Hour(),
		timeNow.Minute(), timeNow.Second())

	// 方式二
	fmt.Println("------方式二------")
	fmt.Printf(timeNow.Format("2006/01/02 15:04:05\n"))
	fmt.Printf(timeNow.Format("2006-01-02\n"))
	fmt.Printf(timeNow.Format("15:04:05\n"))

}
