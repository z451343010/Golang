package main

// 引入包名时，包的路径从 GO_PATH开始
// 实际的完整路径为：F:\Golang\goProgramProject\src\packageTest\utils
// 编译器自动加入 src
// GO_PATH = F:\Golang\goProgramProject
import (
	"fmt"
	"packageTest/utils"
)

// 在main函数中调用其它包中的函数
func main() {

	var num1 float64
	var num2 float64
	var result float64
	var operator string

	fmt.Println("请输入第一个数：")
	fmt.Scanln(&num1)
	fmt.Println("请输入第二个数：")
	fmt.Scanln(&num2)
	fmt.Println("请输入运算符：")
	fmt.Scanln(&operator)

	result = utils.Calculate(num1, num2, operator)

	fmt.Println("计算的结果为：", result)

}
