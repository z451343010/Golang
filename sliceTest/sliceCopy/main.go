package main

import "fmt"

// 切片的拷贝操作
// 切片使用内置函数 copy 完成拷贝
func main() {

	var slice []int = []int{1, 2, 3, 4, 5}

	var slice2 = make([]int, 10)
	fmt.Println(slice2)

	copy(slice2, slice)
	fmt.Println(slice2)

}
