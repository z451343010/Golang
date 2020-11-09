/*
	冒泡排序
*/
package main

import "fmt"

func bubbleSort(slice []int) []int {

	for j := len(slice) - 1; j > 0; j-- {

		for i := 0; i < j; i++ {

			if slice[i] > slice[i+1] {

				var temp int
				temp = slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = temp

			}

		}

	}

	return slice

}

func main() {

	var slice []int = []int{6, 3, 2, 7, 3, 8, 9}
	fmt.Println(bubbleSort(slice))

}
