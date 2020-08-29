package main

import "fmt"

/*
	类型断言的其它案例
*/

func main() {

	var x interface{}
	var b float32 = 1.1
	x = b // 空接口，可以接收任意类型

	// 使用类型断言
	y := x.(float32)
	fmt.Println(y)

	// 带检测的断言
	z, flag := x.(float64)

	if flag == true {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println(z)

}
