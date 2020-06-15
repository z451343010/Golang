package main

import "fmt"

// 定义二维数组，用于保存三个班，每个班五名同学的成绩
// 并求出每个班级平均分、以及所有班级平均分
func main() {

	var scoreArr [3][5]float64
	var totalSum float64

	for i := 0; i < len(scoreArr); i++ {

		var classSumScore float64
		var classAverageScore float64

		for j := 0; j < len(scoreArr[i]); j++ {

			fmt.Println("请输入第", i+1, "个班，第",
				j+1, "名同学的成绩：")
			fmt.Scanln(&scoreArr[i][j])

			classSumScore += scoreArr[i][j]

		}

		classAverageScore = classSumScore / float64(len(scoreArr[i]))
		totalSum += classAverageScore

		fmt.Println("第", i+1, "班的平均成绩为：", classAverageScore)

	}

	fmt.Println("所有班级平均分为：", totalSum/float64(len(scoreArr)))

}
