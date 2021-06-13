/*
	牛客网
	华为机试
	HJ16
*/
package main

import "fmt"

func maxNum(compareNum []int) int {
	max := compareNum[0]
	for _, value := range compareNum {

		if value >= max {
			max = value
		}

	}

	return max
}

func main() {

	var N, m int
	fmt.Scanf("%d %d\n", &N, &m)

	dp := make([]int, 32000)
	// 主件1 的价格
	zj := make([]int, m+1)
	// 主件1 的价格*重要度
	zjvw := make([]int, m+1)
	// 附件1，附件2的价格
	fj1 := make([]int, m+1)
	fj2 := make([]int, m+1)
	// 附件1，附件2的价格*重要度
	fj1vw := make([]int, m+1)
	fj2vw := make([]int, m+1)

	for i := 1; i <= m; i++ {
		var v, p, q int
		fmt.Scanf("%d %d %d\n", &v, &p, &q)
		// 主件
		if q == 0 {
			zj[i] = v
			zjvw[i] = v * p
		} else if fj1[q] == 0 { // 附件1
			fj1[q] = v
			fj1vw[q] = v * p
		} else if fj2[q] == 0 {
			fj2[q] = v
			fj2vw[q] = v * p
		}

	}

	// 动态规划
	for i := 1; i <= m; i++ {
		for j := N; j >= 0; j-- {
			// 只要主件
			if j >= zj[i] {
				numSlice := make([]int, 2)
				numSlice[0] = dp[j]
				numSlice[1] = dp[j-zj[i]] + zjvw[i]
				dp[j] = maxNum(numSlice)
			}
			// 主件 + 附件1
			if j >= zj[i]+fj1[i] {
				numSlice := make([]int, 2)
				numSlice[0] = dp[j]
				numSlice[1] = dp[j-zj[i]-fj1[i]] + zjvw[i] + fj1vw[i]
				dp[j] = maxNum(numSlice)
			}
			// 主件 + 附件2
			if j >= zj[i]+fj2[i] {
				numSlice := make([]int, 2)
				numSlice[0] = dp[j]
				numSlice[1] = dp[j-zj[i]-fj2[i]] + zjvw[i] + fj2vw[i]
				dp[j] = maxNum(numSlice)
			}
			// 主件 + 附件1 + 附件2
			if j >= zj[i]+fj1[i]+fj2[i] {
				numSlice := make([]int, 2)
				numSlice[0] = dp[j]
				numSlice[1] = dp[j-zj[i]-fj1[i]-fj2[i]] + zjvw[i] + fj1vw[i] + fj2vw[i]
				dp[j] = maxNum(numSlice)
			}
		}
	}

	fmt.Println(dp[N])

}
