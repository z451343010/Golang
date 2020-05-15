package main

import "fmt"

// 对学生成绩大于60分的，输出“合格”。
// 低于60分的，输出“不合格”
// 输入的成绩不能大于100分
func main() {

	var score int

	fmt.Println("请输入学生的成绩：")
	fmt.Scanln(&score)

	switch {
	case score < 60:
		fmt.Println("不合格")
	case score >= 60 && score < 100:
		fmt.Println("合格")
	default:
		fmt.Println("成绩无效")

	}

}
