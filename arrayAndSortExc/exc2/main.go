package main

import "fmt"

// 已知有个排序好（升序）的数组，要求插入一个元素
// 最后打印该数组，顺序依然是升序
func main() {

	var arr = [6]int{1, 2, 4, 5, 7, 8}
	var arr2 = [7]int{0, 0, 0, 0, 0, 0, 0}

	var insertNum int
	fmt.Println("请插入元素")
	fmt.Scanln(&insertNum)

	for i := 0; i < len(arr); i++ {

		if insertNum >= arr[i] {

			arr2[i] = arr[i]

			if i == len(arr)-1 {
				arr2[i+1] = insertNum
			}

		} else {

			arr2[i] = insertNum

			for k := i + 1; k < len(arr2); k++ {
				arr2[k] = arr[k-1]
			}

			break

		}

	}

	fmt.Println(arr2)

}
