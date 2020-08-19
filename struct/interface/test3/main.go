package main

import "fmt"

/*
	只要是自定义数据类型
	就可以实现接口，不仅仅是结构体类型
*/

type integer int

type AInteger interface {
	Say()
}

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}

func main() {
	var i integer = 10
	var b AInteger = i
	b.Say()
}
