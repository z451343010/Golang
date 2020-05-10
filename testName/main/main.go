package main

import (
	"fmt"
	// 为了使用utils.go，需要引入utils.go的包：model
	// 包的绝对路径是：F:\Golang\goProgramProject\src\testName\model
	// 由于我们已经配置好了GOPATH = F:\Golang\goProgramProject
	// src则是编译器自动默认加的
	"testName/model"
)

// 如果变量名、函数名、常量名首字母大写，则可以被其它的包访问
// 如果首字母小写，则只能在本包中使用
// 可以理解成，首字母大写是公有的，首字母小写是私有的
// 在Go语言中，没有 public private 等关键字
func main() {
	// 使用 utils.go里面的变量
	fmt.Println(model.StuName)
}
