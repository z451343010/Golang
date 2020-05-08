package main

import "fmt"

// Golang中基本数据类型的转换
func main() {
	var i int = 100
	var n1 float32 = float32(i)
	fmt.Println("i = ", i, " n1 = ", n1)

	// 很多语言中，低精度到高精度的转换可以自动转换
	// 但是在Golang中，必须程序员自己强制转换
	var n3 int64 = int64(i)
	fmt.Println("n3 = ", n3)

	// 范围大的数据类型转换为范围小的数据类型
	var theNum int64 = 25562
	var theNum2 int8 = int8(theNum)
	fmt.Println(" theNum2 = ", theNum2)

}
