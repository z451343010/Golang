/*
	快速排序
*/
package main

import "fmt"

func quickSort(slice []int, left int, right int) {

	if left >= right {
		return
	}

	// 默认第一个数为基准数
	var pivot int = slice[left]
	// 从第二个数开始扫描
	var i int = left + 1
	var j int = right

	for i < j {

		for slice[i] < pivot {
			i++
		}

		for slice[j] > pivot {
			j--
		}

		if i < j {
			slice[i], slice[j] = slice[j], slice[i]
		}

	}

	var pivotIndex = j
	if slice[left] > slice[pivotIndex] {
		slice[left], slice[pivotIndex] = slice[pivotIndex], slice[left]
	}

	quickSort(slice, left, pivotIndex-1)
	quickSort(slice, pivotIndex+1, right)

}

func main() {
	var slice []int = []int{108, 3, 88, 12, 8, 88, 555, 7, 91, 35, 66}
	quickSort(slice, 0, len(slice)-1)
	fmt.Println("slice = ", slice)
}
