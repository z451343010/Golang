package main

import "fmt"

// 变量的作用域
func test() {

	// age和Name的作用域就只在函数内部
	var age int = 10
	var Name string = "tom"
	fmt.Println(age)
	fmt.Println(Name)

}

func main() {

	// fmt.Println(age)  main函数中不能调用test里面的局部变量

}
