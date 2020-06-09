package main

import "fmt"

// 如果想在其它函数中
// 去修改原来的数组，可以使用 （指针方式）
func test(testArr *[3]int) {
	(*testArr)[0] = 88
}

func main() {
	var arr = [3]int{11, 22, 33}
	test(&arr)
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr)
	var ptr *[3]int = &arr
	fmt.Println(&ptr)
	fmt.Println(&arr[0])
	// fmt.Println(&arr == &arr[0])
}
