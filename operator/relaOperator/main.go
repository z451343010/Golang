package main

import "fmt"

var num int = 5
var num2 int = 6

// 关系运算符的演示
func main() {
	fmt.Println(num == num2)
	fmt.Println(num != num2)
	fmt.Println(num >= num2)
	fmt.Println(num <= num2)
	fmt.Println(num < num2)
	fmt.Println(num > num2)

	// 把关系运算符的结果赋值给一个 bool 类型
	var flag bool

	flag = num == num2
	fmt.Println("flag = ", flag)

	flag = num < num2
	fmt.Println("flag = ", flag)

}
