package main

import "fmt"

// 通常在init函数中完成一些初始化的工作
func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("main function")
}
