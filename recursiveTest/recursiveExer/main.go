package main

import "fmt"

// 求函数值
// 已知f(1)=3;f(n)=2*f(n-1)+1;
// 请使用递归的思想编程，求出f(n)的值
func f(n int) int {

	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}

}

func main() {

	var n int
	fmt.Println("请输入n")
	fmt.Scanln(&n)

	fmt.Println(f(n))

}
