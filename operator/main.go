package main

import "fmt"

// 求两个数的最大值
func main() {

	var num1 int = 20
	var num2 int = 80
	var maxNum int

	if num1 >= num2 {
		maxNum = num1
	} else {
		maxNum = num2
	}

	fmt.Println("maxNum = ", maxNum)

}
