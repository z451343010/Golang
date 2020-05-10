package main

import "fmt"

// 运算符
func main() {

	// 除法的使用
	// 如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4) // 原来应该是2.5，去掉小数部分后成为2

	var testNum float32 = 10 / 4
	fmt.Println(testNum) // 结果仍然是2

	//如果我们希望保留小数部分，则需要有浮点数参与运算
	var testNum2 float32 = float32(10) / float32(4)
	fmt.Println(testNum2)

	// 取模（%）的使用
	// a % b = a - a / b *b
	fmt.Println(-10 % 3) // =-10-(-10)/3*3 = -10 - (-9) = -1

	// ++ 和 -- 的使用
	var i int = 10
	i++            // 等价于 i = i + 1
	fmt.Println(i) // 11
	i--            // 等价于 i = i - 1
	fmt.Println(i) // 10

	// ++ 和 -- 的使用细节
	// 只能独立使用
	var j int = 0
	var k int
	j++
	k = j
	fmt.Println(k)

}
