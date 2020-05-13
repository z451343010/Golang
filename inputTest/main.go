package main

import "fmt"

// Golang中键盘读取输入
func main() {

	// fmt.Scanln()
	var name string
	var age byte
	var salary float32
	var isPass bool

	// fmt.Scanf()
	fmt.Println("请输入你的姓名、年龄、薪水、是否通过。使用空格间隔")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("名字是：%v 年龄是：%v 薪水是：%v 是否通过：%v ", name, age, salary, isPass)

}
