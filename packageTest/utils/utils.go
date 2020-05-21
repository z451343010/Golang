package utils

import "fmt"

// 定义一个计算函数
// 需要挎包调用时，函数名需要大写
// 类似于Java的public
// 在Golang里面称为该函数可导出
func Calculate(num1 float64, num2 float64, operator string) float64 {

	var result float64

	switch operator {

	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("输入的结果有误")

	}

	return result

}
