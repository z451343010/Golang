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

		if inputNum == 0 {
			break
		}

		if inputNum > 0 {
			positiveNum++
			// 当 inputNum > 0 时，往下执行已经无意义
			// 因此可以用continue跳出循环
			// 不再执行下面的语句
			continue
		}

		// 只有当输入的数字为负数时，才能执行到这条语句
		// 因为输入的数字为0时已经跳出循环
		// 输入的数字为正数的时候，跳过本次循环
		negativeNum++

	}

	fmt.Println("输入的正数个数为：", positiveNum)
	fmt.Println("输入的负数个数为：", negativeNum)

}
