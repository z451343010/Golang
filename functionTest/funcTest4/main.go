package main

import "fmt"

// 函数既然是一种数据类型。
// 那么在Go中，函数可以作为形参，并且调用。

func testFun(num1 int, num2 int) int {

	return num1 + num2

}

// testFun func(int, int) int 便是作为传入的形参函数
func getSum(testFun func(int, int) int, num1 int, num2 int) int {

	return testFun(num1, num2)

}

func main() {

	res := getSum(testFun, 50, 60)
	fmt.Printf("res的数据类型是：%T ", res)

	fmt.Println("res的值是：", res)

	res2 := getSum(testFun, 50, 60)
	fmt.Println("res2 =", res2)

}
