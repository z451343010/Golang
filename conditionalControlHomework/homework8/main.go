package main

import "fmt"

// 判断考试成绩为什么等级
// 90-100：优秀
// 80-89：优良
// 70-79：良好
// 60-69：及格
// 60以下：不及格
func main() {

	var score float64
	fmt.Println("请输入学生的成绩：")
	fmt.Scanln(&score)

	// 使用类似于 if else 的用法时
	// switch 后面不需要加表达式

	switch {

	case score >= 90 && score <= 100:
		fmt.Println("优秀")
	case score >= 80 && score < 90:
		fmt.Println("优良")
	case score >= 70 && score < 80:
		fmt.Println("良好")
	case score >= 60 && score < 70:
		fmt.Println("及格")
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	default:
		fmt.Println("输入的成绩有误")
	}

}
