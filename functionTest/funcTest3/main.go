package main

import "fmt"

// 在Go中，函数也是一种数据类型
// 可以赋值给一个变量，则该变量就是一个函数类型的变量了
// 通过该变量可以对函数调用

func getSum(n1 int, n2 int) int {

	return n1 + n2

}

func main() {

	a := getSum
	fmt.Printf("a的类型是%T,getSum类型是%T", a, getSum)

}
