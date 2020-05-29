package main

import "fmt"

// 如果将匿名函数赋值给一个全局变量
// 那么这个匿名函数，就成为一个全局匿名函数
// 可以在程序有效。
var (

	// func1就是一个全局匿名函数
	func1 = func(num1 int, num2 int) int {
		return num1 + num2
	}
)

func main() {

	result := func1(10, 20)
	fmt.Println("result = ", result)

}
