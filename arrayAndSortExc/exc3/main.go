package main

import "fmt"

// 保存 1 3 5 7 9 五个技术到数组
// 并倒序打印
func main() {

	var numArr [5]int

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			numArr[i/2] = i
		}
	}

	fmt.Println(numArr)
	fmt.Println()
	fmt.Println("倒序打印数组：")

	// 冒泡排序(倒序)
	for i := len(numArr) - 1; i > 0; i-- {

		for j := 0; j < i; j++ {
			if numArr[j] < numArr[j+1] {
				var tempNum int
				tempNum = numArr[j]
				numArr[j] = numArr[j+1]
				numArr[j+1] = tempNum
			}
		}

	}

	fmt.Println(numArr)

}
