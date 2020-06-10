package main

import (
	"fmt"
	"math"
)

// leetCode练习题
// 判断一个整数是否是回文数
// 回文数是指正序（从左向右）和倒序（从右向左）都是一样的数
func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var remainSlice []int
	var tempX int = x
	// 取出整数所有位上面的数字
	for {

		var remain int
		remain = x % 10
		x = int(x / 10)
		remainSlice = append(remainSlice, remain)

		if x == 0 {
			break
		}

	}

	var y float64
	for i := 0; i < len(remainSlice); i++ {
		y += float64(remainSlice[i]) * math.Pow(10,
			float64(len(remainSlice))-1-float64(i))
	}

	if int(y) == tempX {
		return true
	} else {
		return false
	}

}

func main() {

	fmt.Println(isPalindrome(-121))

}
