package main

import "fmt"

// 编写一个函数，输出100以内的所有素数
// 每行显示5个，并求和
func getPrimeSum(n int) (sum int) {

	var count int

label1:
	for i := 2; i <= n; i++ {

		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue label1
			}
		}

		count++
		fmt.Print(i, "\t")
		if count%5 == 0 {
			fmt.Println()
		}

		sum += i

	}

	return sum

}

func main() {

	var n int = 100
	var sum int = getPrimeSum(n)
	fmt.Println("100以内质数的和为：")
	fmt.Println(sum)

}
