package main

import "fmt"

// 请求出一个数组的最大值，并得到对应的下标
func main() {

	var numArr = [5]float64{9.9, 3.5, 110.72, 3.3, 56}
	var arrMaxNum float64
	var arrMaxIndex int

	for index, value := range numArr {
		if value > arrMaxNum {
			arrMaxNum = value
			arrMaxIndex = index
		}
	}

	fmt.Println(arrMaxNum)
	fmt.Println(arrMaxIndex)

}
