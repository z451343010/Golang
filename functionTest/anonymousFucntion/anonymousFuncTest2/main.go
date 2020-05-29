package main

import "fmt"

// 匿名函数使用方式2：
// 将匿名函数赋值给一个变量（函数变量）
// 再通过该变量来调用匿名函数
func main() {

	// testFunc相当于一个函数类型的变量
	testFunc := func(num1 int, num2 int) int {

		return num1 + num2

	}

	var result int = testFunc(10, 20)
	fmt.Println("result = ", result)

}
