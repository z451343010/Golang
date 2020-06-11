package main

import "fmt"

// slice定义方式三
// 定义一个切片，直接就指定具体数组
// 使用原理类似make的方式
func main() {

	var slice []int = []int{1, 2, 3}
	fmt.Println(slice)

}
