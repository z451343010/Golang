/*
	插入排序
	适用于基本有序的数组
	思路，从第二个元素开始，和前面的元素不断比较
	如果小于前面的元素，则和前面的元素对换，往前挪动
*/
package main

import "fmt"

func insertSort(slice []int) {

	for i := 1; i < len(slice); i++ {

		for j := i; j > 0; j-- {

			if slice[j] < slice[j-1] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}

		}

	}

}

func main() {

	var slice []int = []int{108, 3, 88, 12, 8, 88, 555, 7, 91, 35, 66}
	insertSort(slice)
	fmt.Println("slice = ", slice)

}
