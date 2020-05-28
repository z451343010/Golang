package main

import "fmt"

// 编写一个函数swap
// 可以交换n1和n2的值
func swap(n1 *int, n2 *int) {

	var temp int = *n1
	*n1 = *n2
	*n2 = temp

}

func main() {

	var n1 int = 2
	var n2 int = 3

	// 如果希望函数内的变量能修改函数外的变量
	// 可以传入变量的地址&
	// 函数内以指针的方式操作变量
	swap(&n1, &n2)

	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)

}
