package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 课堂练习
// 编写一个函数，随机猜数游戏
// 随机生成一个1-100的整数
// 有十次机会
func main() {

	rand.Seed(time.Now().UnixNano())
	var randNum int = rand.Intn(100) + 1
	fmt.Println(randNum)

	var guessTimes int
	var guessNum int
	for i := 0; i < 10; i++ {

		guessTimes++
		fmt.Println("请输入随机数：")
		fmt.Scanln(&guessNum)
		if guessNum == randNum {
			break
		}

	}

	switch guessTimes {
	case 1:
		fmt.Println("你真是个天才")
	case 2, 3:
		fmt.Println("你很聪明，快赶上我了")
	case 4, 5, 6, 7, 8, 9:
		fmt.Println("一般般")
	case 10:
		if guessNum == randNum {
			fmt.Println("可算猜对啦")
		} else {
			fmt.Println("说你点啥好呢")
		}

	}

}
