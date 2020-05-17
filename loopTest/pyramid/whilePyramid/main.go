package main

import "fmt"

// 打印三角形金字塔，使用while方式
func main() {

	var layersNum int
	fmt.Println("请输入层数：")
	fmt.Scanln(&layersNum)

	var i int = 1
	for {

		var j int = 1
		for {

			if j > layersNum-i {
				break
			}

			fmt.Print(" ")
			j++

		}

		var k int = 1
		for {

			if k > 2*i-1 {
				break
			}

			fmt.Print("*")
			k++

		}

		i++
		if i > layersNum {
			break
		}
		fmt.Println()

	}

}
