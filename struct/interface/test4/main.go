package main

import "fmt"

/*
	一个接口（比如A接口）可以继承多个别的接口（比如B,C接口），
	这时如果要实现A接口，也必须将B,C接口的方法也全部实现。
*/

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Test struct {
}

func (t Test) Hello() {
	fmt.Println("Test interface Hello")
}

func (t Test) Say() {
	fmt.Println("Test interface Say()")
}

func main() {

	var test Test
	var a AInterface = test
	var b BInterface = test
	a.Say()
	b.Hello()

}
