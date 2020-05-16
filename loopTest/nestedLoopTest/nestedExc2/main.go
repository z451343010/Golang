package main

import "fmt"

// 统计三个班级及格人数
// 每个班级有五名同学
func main() {

	for i := 1; i <= 3; i++ {

		var classPassNum int

		if i > 1 {
			fmt.Println()
		}

		fmt.Println("请输入第", i, "个班级的学生的成绩：")

		for j := 1; j <= 5; j++ {

			var score float64
			fmt.Scanln(&score)
			if score >= 60 {
				classPassNum++
			}

		}

		fmt.Println("第", i, "个班级的及格人数为：", classPassNum)

	}

}
