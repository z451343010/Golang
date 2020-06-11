package main

import "fmt"

// 演示切片的基本使用
func main() {

	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// slice表示引用到intArr这个数组的第2个（下标为1）
	// 到下标为3的元素（不包含3）
	slice := intArr[1:3]
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice = ", slice)
	fmt.Println("slice 的容量 = ", cap(slice))

	slice[1] = 34
	fmt.Println(intArr)

}
