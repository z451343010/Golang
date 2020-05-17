package main

import (
	"fmt"
	"math"
)

// 打印一个n层的菱形
func main() {

	var layersNum int
	fmt.Println("请输入要打印的菱形的层数：")
	fmt.Scanln(&layersNum)

	if layersNum%2 == 0.0 {
		fmt.Println("菱形的层数必须为单数")
		return
	}

	var i float64 = 1
	for ; i <= float64(layersNum); i++ {

		// 打印空格
		var j float64 = 1
		for ; j <= math.Abs(i-(float64(layersNum)+1)/2); j++ {
			fmt.Print(" ")
		}

		// 打印*
		if i <= (float64(layersNum)+1)/2 { // 打印上三角

			var k float64 = 1
			for ; k <= 2*i-1; k++ {
				if k == 1 || k == 2*i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}

		} else {
			// 打印下三角
			var l float64 = 1
			for ; l <= (float64(layersNum)+1-i)*2-1; l++ {
				if l == 1 || l == (float64(layersNum)+1-i)*2-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}

		}

		fmt.Println()

	}

}
