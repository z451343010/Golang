package main

import "fmt"

// 打印 1~100 之间所有是 9 的倍数的整数的个数及总和
func main() {

	var count byte
	var sum int

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}

	fmt.Println("count = ", count)
	fmt.Println("sum = ", sum)

}
