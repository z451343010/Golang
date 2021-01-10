/*
	排序
	冒泡排序
*/
package main

import "fmt"

func bubbleSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {

		for j := 0; j < len(slice)-1-i; j++ {

			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}

		}

	}

}

func main() {
	var slice []int = []int{108, 3, 88, 12, 8,
		88, 555, 7, 91, 35, 66}
	// slice 切片为引用传递
	// 因此，经过排序后，slice 的内容会改变
	bubbleSort(slice)
	fmt.Println("sorted slice = ", slice)
}
