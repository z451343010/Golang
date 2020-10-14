/*
	goroutine中使用recover
	解决协程中出现panic，导致程序崩溃问题
*/
package main

import (
	"fmt"
	"time"
)

// 函数
func sayHello() {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}

}

// 函数
func test() {

	defer func() {
		// 捕获 test 抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误，错误原因为：", err)
		}

	}()

	// 定义了一个map
	var myMap map[int]string
	myMap[0] = "Golang" // error,因为没有make()

}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ", i)
		time.Sleep(time.Second)
	}

}
