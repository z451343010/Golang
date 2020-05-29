package main

import "fmt"

// 如果一个文件同时包含全局变量定义，init函数和main函数
// 则执行的流程是：全局变量定义 ->  init函数 -> main函数
func init() {

	fmt.Println("init")

}

// 为了更直观反应函数的执行顺序
// 这里的变量定义为一个函数变量
var age = test()

func test() int {

	fmt.Println("age test()")
	return 90

}

func main() {

	fmt.Println("main function")
	// 由于定义的age是全局变量
	// 因此可以直接在main函数当中调用
	fmt.Println("age = ", age)

}
