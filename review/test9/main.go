/*
	复习
	递归
*/
package main

import "fmt"

func test(n int) {

	if n > 2 {
		fmt.Println(n)
		return
	} else {
		n++
		test(n)
	}

}

func test2(n int) {

	if n > 2 {
		n--
		test2(n)
	}

	fmt.Println(n)

}

func main() {

	test(0)
	test2(4)

}
