package main

import "fmt"

// 切片的遍历方式
func main() {

	// 使用常规的for循环遍历
	var testSlice []int = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(testSlice); i++ {
		fmt.Println(testSlice[i])
	}

	// 使用 for range
	fmt.Println("使用for range遍历：")
	for _, value := range testSlice {
		fmt.Println(value)
	}

}
