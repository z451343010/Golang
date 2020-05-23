package main

import "fmt"

// 给出一个整数n
// 求出它的斐波那契数列
func fibonacci(n int) int {

	var fibonacciNum int

	if n <= 2 {
		// fmt.Println(1)
		return 1
	} else {
		fibonacciNum = fibonacci(n-1) + fibonacci(n-2)
		// fmt.Println(fibonacciNum)
		return fibonacciNum
	}

}

func main() {

	var n int
	fmt.Println("请输入n值")
	fmt.Scanln(&n)

	fmt.Println(fibonacci(n))

}
