package main

import "fmt"

// 指针运算符，取地址，去值
func main() {
	var a int = 100
	fmt.Println("the address of a is:", &a)
	var ptr *int = &a
	fmt.Println("the value of ptr is:", ptr)
	fmt.Println("the value which address is:", *ptr)

}
