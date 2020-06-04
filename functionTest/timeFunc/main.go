package main

import (
	"fmt"
	"time"
)

func main() {

	// 显示系统当前的时间
	timeNow := time.Now()
	fmt.Printf("the value of time is %v\n", timeNow)
	fmt.Printf("the type of time is %T\n", timeNow)

	// 通过time.Now()同样可以获取到年月日时分秒
	fmt.Println("年=", timeNow.Year())
	fmt.Println("月=", timeNow.Month())
	fmt.Println("日=", timeNow.Day())
	fmt.Println("时=", timeNow.Hour())
	fmt.Println("分=", timeNow.Minute())
	fmt.Println("秒=", timeNow.Second())

}
