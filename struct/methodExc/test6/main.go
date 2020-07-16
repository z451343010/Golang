package main

import "fmt"

/*
	实现小小计算器结构体（Calculator）
	实现加减乘除四个功能
*/
type Calculator struct {
}

func (cal *Calculator) calculate(num1 float64, num2 float64,
	operator byte) (result float64) {

	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	default:
		fmt.Println("输入有误")
	}

	return result

}

func main() {

	var num1 float64
	var num2 float64
	var operator byte

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanf("%c", &operator)

	var calculator Calculator
	result := (&calculator).calculate(num1, num2, operator)

	fmt.Println(result)

}
