package main

import "fmt"

// 跳水比赛，8个评委打分。运动员的成绩是8个成绩
// 去掉一个最高分，一个最低分，剩下的6个分数的评委就是
// 最后得分。使用以为数组实现以下功能：
// （1）请把最高分的评委和最低分的评委找出来。
// （2）找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委
// 最差评委就是最后得分相差最大的
func main() {

	var scoreArr [8]float64
	var maxScoreMan []int
	var minScoreMan []int
	var sumScore float64

	for i := 0; i < len(scoreArr); i++ {

		fmt.Println("第", i, "号评委打分：")
		var score float64
		fmt.Scanln(&score)
		scoreArr[i] = score
		sumScore += score

	}

	fmt.Println(scoreArr)

	var maxScore float64 = scoreArr[0]
	var minScore float64 = scoreArr[0]

	for i := 1; i < len(scoreArr); i++ {

		if scoreArr[i] > maxScore {
			maxScore = scoreArr[i]
		}

		if scoreArr[i] < minScore {
			minScore = scoreArr[i]
		}

	}

	fmt.Println("最高分为：", maxScore)
	fmt.Println("最低分为：", minScore)

	// 评分最高的评委和评分最低的评委
	for i := 0; i < len(scoreArr); i++ {

		if scoreArr[i] == maxScore {
			maxScoreMan = append(maxScoreMan, i)
		}

		if scoreArr[i] == minScore {
			minScoreMan = append(minScoreMan, i)
		}

	}

	fmt.Println("评分最高的评委为：", maxScoreMan)
	fmt.Println("评分最低的评委为：", minScoreMan)

	var averageScore float64
	averageScore = (sumScore - maxScore - minScore) /
		(float64(len(scoreArr)) - 2)

	fmt.Println("该运动员的最后得分是：", averageScore)

	var bestMan []int  // 最佳评委
	var worstMan []int // 最差评委

	// 最佳评分
	var bestScore float64 = scoreArr[0] - averageScore
	// 最差评分
	var worstScore float64 = scoreArr[0] - averageScore

	if bestScore < 0 {
		bestScore = -bestScore
	}

	if worstScore < 0 {
		worstScore = -worstScore
	}

	// 统计每个评委评分和运动员最后得分的差
	for i := 1; i < len(scoreArr); i++ {

		var subScore float64 = scoreArr[i] - averageScore

		if subScore < 0 {
			subScore = -subScore
		}

		if subScore < bestScore {
			bestScore = subScore
		}

		if subScore > worstScore {
			worstScore = subScore
		}

	}

	// fmt.Println("最佳评分差是：", bestScore)
	// fmt.Println("最差评分差是：", worstScore)

	// 查找最佳评委、最差评委
	for i := 0; i < len(scoreArr); i++ {

		var subScore float64 = scoreArr[i] - averageScore

		if subScore < 0 {
			subScore = -subScore
		}

		if subScore == bestScore {
			bestMan = append(bestMan, i)
		}
		if subScore == worstScore {
			worstMan = append(worstMan, i)
		}
	}

	fmt.Println("最佳评委为：", bestMan)
	fmt.Println("最差评委为：", worstMan)

}
