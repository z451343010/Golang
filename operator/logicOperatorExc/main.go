package main

import "fmt"

// 有两个变量 a b ，不允许使用中间变量，最终打印结果
func main() {

	var a int = 10
	var b int = 20

	a = a + b
	b = a - b // a + b - b ==> b = a
	a = a - b // a + b - a ==>  a = b

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
