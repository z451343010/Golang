package main

import "fmt"

// Golang中的指针使用
func main() {
	var num int = 3 // 定义一个值为3的int类型变量，变量名为num

	// 1.ptr 是一个指针变量
	// 2.ptr 这个变量的类型是 *int
	// 3.ptr 本身的值是 &num num的地址
	var ptr *int = &num // 把num取地址后，赋值给指针类型的变量 ptr

	fmt.Println("the value of ptr is:", ptr)                                 // 输出指针类型的变量 ptr 的值
	fmt.Println("the value of which it's address is value of ptr is:", *ptr) // 输出 ptr 的值的地址里面存储的内容
	fmt.Println("the address of ptr is:", &ptr)                              // 输出 ptr 指针变量的地址
}
