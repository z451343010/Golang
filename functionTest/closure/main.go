package main

import "fmt"

// 闭包就是一个函数与其相关的引用环境组合的一个整体（实体）
// 累加器
func AddUper() func(int) int {

	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}

}

func main() {

	f := AddUper()
	fmt.Println(f(1))
	fmt.Println(f(2))

}
