/*
	Leetcode 题目 7
	整数反转
	给出一个 32 位的有符号整数
	你需要将这个整数中每位上的数字进行反转。
	ps:
	假设我们的环境只能存储得下 32 位的有符号整数，
	则其数值范围为 [−231,  231 − 1]。
	请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	var numSlice []int
	var result int

	for {
		var quotient int = x / 10
		var remain int = x % 10

		x = quotient
		numSlice = append(numSlice, remain)
		if quotient == 0 {
			break
		}
	}

	for i := 0; i < len(numSlice); i++ {
		var temp int = len(numSlice) - 1 - i
		result += (int(math.Pow10(temp)) * numSlice[i])
	}

	if (result > 2147483647) || (result < -2147483648) {
		return 0
	}

	return result

}

func main() {
	fmt.Println(reverse(123))
}
