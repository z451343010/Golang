/*
	数据结构
	递归
	打印问题
*/
package main

import "fmt"

func test(n int) {

	if n > 2 {
		n--
		test(n)
	}

	fmt.Println("n = ", n)

}

func test2(n int) {

	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n = ", n)
	}

}

func main() {

	test(4)
	fmt.Println("========== test2 ==========")
	test2(4)

}
