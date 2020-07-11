package main

import "fmt"

// 在Golang中的方法作用在指定的数据类型上的
// 和指定的数据类型绑定
// 因此，自定义类型，都可以有方法，而不仅仅是struct
// 比如 int,float32 等都可以有此方法

type integer int

// 将自定义类型 integer 与方法 test() 绑定
func (testInteger integer) test() {

	fmt.Println("testInteger = ", testInteger)

}

// 编写一个方法，可以改变 integer 类型的某个变量的值
func (testInteger2 *integer) test2() {

	(*testInteger2) = 12
	fmt.Println("test2() testInteger2 = ", (*testInteger2))

}

func main() {

	var i integer = 10
	i.test()
	var testInteger2 integer = 11
	(&testInteger2).test2()
	fmt.Println("main testInteger2 = ", testInteger2)

}
