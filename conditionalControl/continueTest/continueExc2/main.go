package main

import "fmt"

// 从键盘读入个数不确定的整数
// 并判断读入的整数和负数的个数
// 输入0时结束程序
// 要求用 for循环 break continue 完成
func main() {

	var positiveNum int
	var negativeNum int

	for {

		var inputNum int
		fmt.Println("请输入数字")
		fmt.Scanln(&inputNum)

		if inputNum > 0 {
			positiveNum++
			continue
		}

		if inputNum < 0 {
			negativeNum++
			continue
		}

		if inputNum == 0 {
			break
		}

	}

	fmt.Println("输入的正数个数为：", positiveNum)
	fmt.Println("输入的负数个数为：", negativeNum)

}
