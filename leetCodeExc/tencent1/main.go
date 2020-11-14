/*
	Leetcode 腾讯算法题1
	给定一个整数数组 nums 和一个目标值 target，
	请你在该数组中找出和为目标值的那两个整数
	并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。
	但是，数组中同一个元素不能使用两遍。
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var result = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				// result = append(result, i)
				// result = append(result, j)
				result[0] = i
				result[1] = j
				return result
			}

		}
	}

	return result

}

func main() {

	var nums []int = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))

}
