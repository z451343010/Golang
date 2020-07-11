package main

import "fmt"

type A struct {
	Num int
}

func (a A) test() {

	a.Num = 12

}

func main() {

	var a A
	a.Num = 13

	// 相当于在调用的时候
	// 把参数 a 传递给 test() 方法的 (a A) 部分
	a.test()
	fmt.Println("main() p.Name = ", a.Num)

}
