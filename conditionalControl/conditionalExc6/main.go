package main

import "fmt"

func main() {

	var score int8 // 分数

	fmt.Println("请输入小明的成绩：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励一辆 BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台 iphone7plus")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一个 ipad")
	} else {
		fmt.Println("什么也没有")
	}

	// 多分支的使用陷阱
	// 当条件1的运算结果为true时，只会执行第一条{}的语句
	var num int8 = 10

	if num > 9 {
		fmt.Println("ok1")
	} else if num > 6 {
		fmt.Println("ok2")
	} else if num > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}

}
