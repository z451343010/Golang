package main

import "fmt"

// defer 延时机制简单案例
func sum(n1 int, n2 int) int {

	// 延时执行
	// 当执行到defer时，暂时不执行
	// 会将defer后面的语句压入到独立的栈（defer栈）
	// 当函数执行完毕，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok1 n2 = ", n2)

	var result int = n1 + n2
	fmt.Println("ok3 result", result)
	return result

}

func main() {

	sum(10, 20)

}
