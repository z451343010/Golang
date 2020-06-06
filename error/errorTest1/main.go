package main

import "fmt"

func test() {

	// 延时函数，先把函数内的匿名错误处理函数压入栈
	defer func() {
		err := recover() // 内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err = ", err)
		}
	}() // 调用匿名函数

	var num int = 10
	var num2 int = 0
	var result int = num / num2
	fmt.Println("result = ", result)

}

// Golang的错误处理机制
func main() {
	test()
}
