package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	Go 协程快速入门
	1）在主线程中（也可以理解成进程）开启一个goroutine
	该协程每隔1秒输出 "hello,world"
	2）在主线程中也每隔1秒输出"hello,golang",输出10次后，退出程序
	3）要求主线程和goroutine同时进行
*/

// 编写一个函数，每隔一秒输出 "hello,world"
func test() {

	for i := 0; i < 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func main() {

	go test() // 开启一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
