package main

import "fmt"

// 变量作用域课堂练习

var name = "tom" // 全局变量

func test01() {
	fmt.Println(name)
}

func test02() {
	name := "jack" // 局部变量
	fmt.Println(name)
}

func main() {

	fmt.Println(name)
	test01()
	test02()
	test01()

}
