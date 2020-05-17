package main

import "fmt"

// 打印空心的三角形金字塔
func main() {

	var layersNum int
	fmt.Println("请输入层数：")
	fmt.Scanln(&layersNum)

	for i := 1; i <= layersNum; i++ {

		for j := 1; j <= layersNum-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}

		fmt.Println()

	}

}
