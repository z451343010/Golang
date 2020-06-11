package main

import "fmt"

func main() {

	// 使用append内置函数
	// 可以对切片进行动态追加
	var slice []int = []int{100, 200, 300}
	slice = append(slice, 400, 500, 600)
	fmt.Println("slice = ", slice)

	fmt.Println("切片后面追加一个切片：")
	slice = append(slice, slice...)
	fmt.Println(slice)

}
