package main

import "fmt"

// 求三个数的最大值，思路2
// 先求出其中两个数的最大值，然后再把那个最大值与第三个数比较
func main() {

	var num1 int = 985
	var num2 int = 211
	var num3 int = 259
	var maxNum int

	// 先求出其中两个数的最大值
	if num1 > num2 {
		maxNum = num1
	} else {
		maxNum = num2
	}

	// 再把 两个数中的最大值 与 第三个数 相比较
	if num3 > maxNum {
		maxNum = num3
	}

	fmt.Println("the max num is ", maxNum)

}
