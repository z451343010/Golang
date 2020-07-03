package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

// 两个结构体具有相同的字段可以互相进行转换
func main() {

	var a A
	var b B

	a = A(b)
	fmt.Println(a, b)

}
