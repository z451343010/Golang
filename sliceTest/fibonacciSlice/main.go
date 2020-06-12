package main

import "fmt"

// 编写一个函数fbn(n int)
// 要求完成：
// 1）可以接收一个n
// 2）能够将斐波那契的数列放到切片中
func fibonacci(num int) int {

	if num == 0 || num == 1 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}

}

func fbn(n int) {

	var fbSlice []int

	for i := 0; i <= n; i++ {
		fbSlice = append(fbSlice, fibonacci(i))
	}

	fmt.Println(fbSlice)

}

func main() {

	fbn(5)

}
