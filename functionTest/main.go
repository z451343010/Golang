// Golang 函数的使用入门案例
package main

import "fmt"

// 将一个计算的功能，放到一个函数中
// 然后在需要使用的时候调用即可
func calculate(num1 float64, num2 float64, operator string) {

	switch operator {
	case "+":
		fmt.Println("函数的计算结果是：", num1+num2)
	case "-":
		fmt.Println("函数的计算结果是：", num1-num2)
	case "*":
		fmt.Println("函数的计算结果是：", num1*num2)
	case "/":
		fmt.Println("函数的计算结果是：", num1/num2)
	default:
		fmt.Println("输入有误")

	}

}

func main() {

	var num1 float64
	var num2 float64
	var operator string

	fmt.Println("请输入第一个数：")
	fmt.Scanln(&num1)
	fmt.Println("请输入第二个数：")
	fmt.Scanln(&num2)
	fmt.Println("请输入运算符：")
	fmt.Scanln(&operator)

	calculate(num1, num2, operator)

}
