package main

import "fmt"

func test(slice []int) {

	slice[0] = 100

}

func main() {

	var slice = []int{1, 2, 3, 4}
	fmt.Println("slice = ", slice)
	test(slice)
	fmt.Println("slice = ", slice)

}
