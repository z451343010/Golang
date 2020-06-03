package main

import "fmt"

// 以函数的方式打印金字塔
// 输入层数，输出对应的金字塔
func pyramid(layerNum int) {

	for i := 0; i < layerNum; i++ {

		for j := 0; j < layerNum-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()

	}

}

func main() {

	fmt.Println("请输入金字塔的层数：")
	var layerNum int
	fmt.Scanln(&layerNum)
	pyramid(layerNum)

}
