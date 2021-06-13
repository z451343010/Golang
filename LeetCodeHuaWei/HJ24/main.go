/*
	牛客网
	华为机试
	HJ24
*/
package main

import (
	"fmt"
	"io"
)

// 重点函数
// 用动态规划来求解递增数列
func returnRankSlice(numSlice []int, rankSlice []int) []int {

	// 初始化递增数列
	// 值都设置为1（因为至少包括自己本身）
	for i := 0; i < len(rankSlice); i++ {
		rankSlice[i] = 1
	}

	for i := 0; i < len(numSlice); i++ {

		var max int
		for j := 0; j < i; j++ {
			if numSlice[j] < numSlice[i] {
				if rankSlice[j] >= max {
					max = rankSlice[j]
				}
			}
		}
		rankSlice[i] = max + 1

	}

	return rankSlice

}

func main() {

	for {

		// 输入人员人数
		var num int
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}

		// 用切片来存储该组人员的身高
		numSlice := make([]int, num)
		for i := 0; i < num; i++ {
			// 循环输入身高
			var height int
			fmt.Scan(&height, " ")
			numSlice[i] = height
		}

		// 求出递增序列
		tempNumSlice1 := make([]int, num)
		incrNumSlice := returnRankSlice(numSlice, tempNumSlice1)

		// 求出反转的 numSlice
		var reNumSlice []int
		for i := len(numSlice) - 1; i >= 0; i-- {
			reNumSlice = append(reNumSlice, numSlice[i])
		}

		// 求出递减序列
		tempNumSlice2 := make([]int, num)
		deNumSlice := returnRankSlice(reNumSlice, tempNumSlice2)

		// 求出反转的递减序列
		var reverseDeNumSlice []int
		for i := len(deNumSlice) - 1; i >= 0; i-- {
			reverseDeNumSlice = append(reverseDeNumSlice, deNumSlice[i])
		}

		// 得到剩下的K位同学最大值
		var maxNum int
		for i := 0; i < len(reverseDeNumSlice); i++ {
			tempNum := reverseDeNumSlice[i] + incrNumSlice[i] - 1
			if tempNum > maxNum {
				maxNum = tempNum
			}
		}

		// 总人数 - 剩下的K位同学的最大值 = 最少需要几位同学出列
		fmt.Println(num - maxNum)
	}

}
