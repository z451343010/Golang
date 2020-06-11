package main

import "fmt"

// 切片是引用类型，因此传递时
// 遵守引用传递机制
func main() {

	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]
	var slice2 = slice
	slice2[0] = 10

	fmt.Println("slice2", slice2)
	fmt.Println("slice", slice)
	fmt.Println("arr", arr)

}
