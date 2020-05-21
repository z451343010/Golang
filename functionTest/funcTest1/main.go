package main

import "fmt"

func test(num1 int) {
	num1 = num1 + 1
	fmt.Println("test() num1 =", num1)
}

func main() {
	var n1 int = 10
	test(n1)
	fmt.Println("main() n1 = ", n1)
}
