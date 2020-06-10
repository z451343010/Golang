package main

import "fmt"

// 请求出一个数组的和以及平均值
func main() {

	var numArr = [5]float64{9.9, 3.5, 110.72, 3.3, 56}
	var arrAverageNum float64
	var arrSum float64

	for _, value := range numArr {

		arrSum += value

	}

	arrAverageNum = arrSum / float64(len(numArr))

	fmt.Println("该数组的和为：", arrSum)
	fmt.Println("该数组的平均值为：", arrAverageNum)

}
