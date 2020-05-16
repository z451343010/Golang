package main

import "fmt"

// Golang 嵌套循环课堂练习1
// 统计3个班级成绩情况，每个班有5个同学
// 求出各个班级的平均分和所有班级的平均分
// (学生成绩从键盘输入)
func main() {

	var totalSum float64     // 所有班级的学生成绩累加
	var totalAverage float64 // 所有班级的平均成绩

	for i := 1; i <= 3; i++ {

		var classSum float64 = 0     // 每个班级的成绩累加
		var classAverage float64 = 0 // 每个班级的平均成绩

		fmt.Println()
		fmt.Println("请输入第", i, "个班级的学生成绩：")

		for j := 1; j <= 5; j++ {
			var score float64
			fmt.Scanln(&score)
			classSum += score
			totalSum += score
		}

		classAverage = classSum / 5

		fmt.Println("第", i, "班的平均成绩为:", classAverage)

	}

	totalAverage = totalSum / (3 * 5)
	fmt.Println("所有班级的平均成绩为：", totalAverage)
}
