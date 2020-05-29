package main

import "fmt"

// 匿名函数使用方式1：
// 在定义匿名函数时就直接调用。这种方式匿名函数只能调用一次
func main() {

	var result int

	result = func(num1 int, num2 int) int {

		return num1 + num2

	}(10, 20) // 直接传入实参调用匿名函数

	fmt.Println("result = ", result)

}
