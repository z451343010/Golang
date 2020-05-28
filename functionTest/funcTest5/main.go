package main

import "fmt"

// 为了简化数据类型定义，Go支持自定义函数类型
func main() {

	// 在Go中，虽然 myInt 和 int 都是 int 类型
	// 但是Go认为，myInt 和 int 不是同一种数据类型
	type myInt int // 给int取了别名

	var num1 myInt
	num1 = 40
	fmt.Println("num1 =", num1)

}
