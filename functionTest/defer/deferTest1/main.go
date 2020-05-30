package main

import "fmt"

// defer 延时机制简单案例
func sum(n1 int, n2 int) int {

	// 延时执行
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok1 n2 = ", n2)

	var result int = n1 + n2
	fmt.Println("ok3 result", result)
	return result

}

func main() {

	sum(10, 20)

}
