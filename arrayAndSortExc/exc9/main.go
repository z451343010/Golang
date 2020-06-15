package main

import "fmt"

// 定义一个数组，并给出8个整数
// 求该数组中大于平均值的数的个数
// 和小于平均值的数的个数
func main() {

	var numArr = [8]float64{1, 11, 4, 9, 6, 3, 4, 8}
	var largerAveNum int
	var lessAveNum int
	var arrSum float64

	for i := 0; i < len(numArr); i++ {

		arrSum += numArr[i]

	}

	for j := 0; j < len(numArr); j++ {

		if numArr[j] > arrSum/float64((len(numArr)-1)) {
			largerAveNum++
		}

		if numArr[j] < arrSum/float64((len(numArr)-1)) {
			lessAveNum++
		}

	}

	fmt.Println("平均值为：", arrSum/float64((len(numArr)-1)))
	fmt.Println("大于平均值的个数为：", largerAveNum)
	fmt.Println("小于平均值的个数为：", lessAveNum)

}
