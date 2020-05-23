package main

import "fmt"

// 经典编程题
// 猴子吃桃
// 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个
// 以后猴子每天都吃其中的一半，然后再多吃一个。
// 当到第10天时，想再吃，发现只有一个桃子了
// 问题：最初有多少个桃子
func monkeyEatPeach(day int) int {

	if day > 10 || day < 1 {

		fmt.Println("输入的桃子数量不对")
		return 0

	}

	if day == 10 {

		// 第十天只有1个桃子
		return 1

	} else {

		// 假设第n天的桃子树为y个
		// 第n-1天的桃子树为x个
		// 根据题意: x/2 - 1 = y
		// 因此，x = 2 * (y + 1)
		return 2 * (monkeyEatPeach(day+1) + 1)

	}

}

func main() {

	fmt.Println("最初的桃子数量为：")
	fmt.Println(monkeyEatPeach(1))

}
