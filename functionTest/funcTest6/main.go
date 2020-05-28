package main

import "fmt"

type myFunc func(int, int) int // 自定义函数类型的变量

func getSum(num1 int, num2 int) int {

	return num1 + num2

}

func testMyFunc(funcVar myFunc, num1 int, num2 int) int {

	return funcVar(num1, num2)

}

func main() {

	// 定义一个 myFunc 类型的变量
	// myFunc 为上面定义的函数类型的变量
	var myFuncVar myFunc
	myFuncVar = getSum
	fmt.Println(testMyFunc(myFuncVar, 10, 20))

}
