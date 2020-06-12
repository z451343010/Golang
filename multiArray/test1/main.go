package main

import "fmt"

// 二维数组演示案例
// 用二维数组输出如下的图形
// 0 0 0 0 0 0
// 0 0 1 0 0 0
// 0 2 0 3 0 0
// 0 0 0 0 0 0

func main() {

	// 定义声明二维数组
	var multiArr [4][6]int // [行][列]

	multiArr[1][2] = 1
	multiArr[2][1] = 2
	multiArr[2][3] = 3

	// 遍历二维数组
	for i := 0; i < len(multiArr); i++ {

		for j := 0; j < len(multiArr[i]); j++ {
			fmt.Print(multiArr[i][j])
		}

		fmt.Println()

	}

}
