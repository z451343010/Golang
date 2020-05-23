package main

import "fmt"

// 函数的递归调用
// 递归：函数调用自己本身
func test(n int) {

	if n > 2 {
		n--
		test(n)
	}

	fmt.Println("n =", n)

}

func main() {

	test(4)

}
