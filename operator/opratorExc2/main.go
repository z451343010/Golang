package main

import "fmt"

// 求出三个数的最大值
func main() {

	var num1 int = 33
	var num2 int = 12
	var num3 int = 126
	var maxNum int

	if num1 >= num2 && num2 >= num3 {
		maxNum = num1
	} else if num1 < num2 && num2 >= num3 {
		maxNum = num2
	} else {
		maxNum = num3
	}

	fmt.Println(maxNum)

}
