package main

import "fmt"

// 二维数组的遍历
func main() {

	// 双重for循环
	var multiArr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(multiArr); i++ {
		for j := 0; j < len(multiArr[i]); j++ {
			fmt.Print(multiArr[i][j], " ")
		}
		fmt.Println()
	}

	// for - range 方式
	for _, value := range multiArr {

		for _, value2 := range value {
			fmt.Print(value2, " ")
		}

		fmt.Println()

	}

}
