package main

import "fmt"

// 从终端输入一个数字
// 打印出对应的乘法表
func multiTable(num int) {

	for i := 1; i <= num; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i, "\t")
		}

		fmt.Println()

	}

}

func main() {

	fmt.Println("请输入数字：")
	var num int
	fmt.Scanln(&num)
	multiTable(num)

}
