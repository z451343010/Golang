/*
	复习
	经典编程题：猴子吃桃
*/
package main

import "fmt"

func monkeyEatPeach(day int) int {

	if day > 10 || day < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}

	if day == 10 {
		return 1
	} else {
		return 2 * (monkeyEatPeach(day+1) + 1)
	}

}

func main() {

	fmt.Println("最初的桃子数量为：")
	fmt.Println(monkeyEatPeach(1))

}
