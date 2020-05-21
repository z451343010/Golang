package main

import "fmt"

// 输入两个数
// 再输入一个运算符(+,-,*,/)
// 得到结果
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

	switch operator {
	case "+":
		fmt.Println("运算结果是：", num1+num2)
	case "-":
		fmt.Println("运算结果是：", num1-num2)
	case "*":
		fmt.Println("运算结果是：", num1*num2)
	case "/":
		fmt.Println("运算结果是：", num1/num2)
	default:
		fmt.Println("输入有误")
	}

}
