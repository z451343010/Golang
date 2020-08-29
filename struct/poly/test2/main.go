package main

import "fmt"

// 类型断言

type Point struct {
	x int
	y int
}

func main() {

	var a interface{} // 空接口
	var point Point = Point{1, 2}
	a = point // 默认实现空接口

	// 如何将 a 赋值给一个 Point 变量
	var b Point

	// 类型断言，判断 a 是否指向 Point 类型的变量
	// 如果是就转成 Point 类型并赋给 b 变量，否则报错
	b = a.(Point)
	fmt.Println(b)

}
