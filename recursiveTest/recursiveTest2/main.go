package main

import "fmt"

// 递归案例2
func test2(n int) {

	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n =", n)
	}

}

func main() {

	test2(4)

}
