package main

import "fmt"

// 编写一个函数，可以接收一个数组，该数组有5个数
// 请找出最大的数和最小的数和对应的数组下标是多少
func main() {

	var numArr = [5]int{2, 5, 6, 3, 89}
	var maxNum int
	var maxNumIndex int
	var minNum int
	var minNumIndex int

	maxNum = numArr[0]
	minNum = numArr[0]

	for i := 1; i < len(numArr); i++ {

		if numArr[i] > maxNum {
			maxNum = numArr[i]
			maxNumIndex = i
		}

		if numArr[i] < minNum {
			minNum = numArr[i]
			minNumIndex = i
		}

	}

	fmt.Println("最大的数为：", maxNum, "最大的数的下标为：", maxNumIndex)
	fmt.Println("最小的数为：", minNum, "最小的数的下标为：", minNumIndex)

}
