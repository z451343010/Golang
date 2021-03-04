/*
	复习
	有两个变量，不使用中间变量
	最终交换两个变量的值，并打印
*/
package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
