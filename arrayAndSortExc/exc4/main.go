package main

import "fmt"

// 定义一个3行4列的二维数组
// 逐个从键盘输入值，编写程序将四周的数据清零0
func main() {

	var numArr [3][4]int

	// 输入数据
	for i := 0; i < len(numArr); i++ {

		for j := 0; j < len(numArr[i]); j++ {
			var num int
			fmt.Scanln(&num)
			numArr[i][j] = num
		}

	}

	fmt.Println()
	fmt.Println("四周清零后的二维数组为")

	// 将程序四周的数据清0
	for i := 0; i < len(numArr); i++ {

		for j := 0; j < len(numArr[i]); j++ {

			if j == 0 || j == (len(numArr[i])-1) {
				numArr[i][j] = 0
			}

			if i == 0 || i == len(numArr)-1 {
				numArr[i][j] = 0
			}
		}

	}

	// 打印后面的数组
	for i := 0; i < len(numArr); i++ {

		for j := 0; j < len(numArr[i]); j++ {
			fmt.Print(numArr[i][j], "\t")
		}

		fmt.Println()

	}

}
