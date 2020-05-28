package main

import "fmt"

// Go支持对函数返回值进行命名
func getSumAndSub(num1 int, num2 int) (sum int, sub int) {

	sum = num1 + num2
	sub = num1 - num2

	return

}

func main() {

	var sub int
	var sum int
	var sub2 int

	// 若只想获取其中一个返回值，则另一个用_省略
	_, sub = getSumAndSub(10, 20)
	sum, sub2 = getSumAndSub(10, 20)

	fmt.Println("sub =", sub)
	fmt.Println("sum =", sum)
	fmt.Println("sub2 =", sub2)

}
