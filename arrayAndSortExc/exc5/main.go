package main

import "fmt"

// 定义一个4*4的二维数组
// 逐个从键盘输入值
// 然后将第1行的第4行的数据进行交换
// 然后将第2行的第3行的数据进行交换

// 输出数组的函数
func outputArr(numArr [4][4]int) {

	for i := 0; i < len(numArr); i++ {

		for j := 0; j < len(numArr[i]); j++ {
			fmt.Print(numArr[i][j], "\t")
		}

		fmt.Println()

	}

	fmt.Println()

}

func main() {

	var numArr [4][4]int

	fmt.Println("请输入数据：")

	// 输入数据
	for i := 0; i < len(numArr); i++ {

		for j := 0; j < len(numArr[i]); j++ {
			var num int
			fmt.Scanln(&num)
			numArr[i][j] = num
		}

	}

	fmt.Println("交换前的数组为：")
	outputArr(numArr)

	var tempNumArr [4][4]int
	for i := 0; i < len(tempNumArr); i++ {

		for j := 0; j < len(tempNumArr[i]); j++ {
			tempNumArr[i][j] = numArr[i][j]
		}

	}

	// 交换
	for i := 0; i < len(numArr); i++ {

		for j := 0; j < len(numArr[i]); j++ {
			numArr[i][j] = tempNumArr[len(tempNumArr)-1-i][j]
		}

	}

	fmt.Println("交换后的数组为：")
	outputArr(numArr)

}
