package main

import "fmt"

// 编写一个函数
// 计算两个数的和、差
// 并返回计算结果
func getSumAndSub(num1 float64, num2 float64) (float64, float64) {
	return num1 + num2, num1 - num2
}

func main() {

	var sumRes float64
	var subRes float64

	var num1 float64
	var num2 float64

	fmt.Println("请输入第一个数：")
	fmt.Scanln(&num1)
	fmt.Println("请输入第二个数：")
	fmt.Scanln(&num2)

	sumRes, subRes = getSumAndSub(num1, num2)
	_, subRes = getSumAndSub(num1, num2)
	fmt.Println("subRes = ", subRes)

	fmt.Println("两个数的和为：", sumRes)
	fmt.Println("两个数的差为：", subRes)

}
