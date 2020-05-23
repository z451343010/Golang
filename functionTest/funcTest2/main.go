package main

import "fmt"

func getSum(num1 float64, num2 float64) float64 {

	// 当函数有return时
	// 就是相当于把函数执行的结果返回给调用者
	return num1 + num2

}

func main() {

	var num1 float64 = 19.36
	var num2 float64 = 20.25

	fmt.Println(getSum(num1, num2))

}
