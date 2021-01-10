/*
	排序
	快速排序
*/
package main

import "fmt"

func quickSort(left int, right int, slice []int) {

	if left >= right {
		return
	}

	pivot := slice[left]
	i := left + 1
	j := right

	for i <= j {

		fmt.Println(" i = ", i)
		fmt.Println(" j = ", j)

		if i < j {
			slice[i], slice[j] = slice[j], slice[i]
		}

		for slice[i] < pivot {
			i++
		}

		for slice[j] > pivot {
			j--
		}

	}

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

	slice[left], slice[j] = slice[j], slice[left]
	// quickSort(left, j-1, slice)
	// quickSort(j+1, right, slice)

}

func main() {

	// var slice []int = []int{108, 3, 88, 12, 8, 88, 555, 7, 91, 35, 66}
	// quickSort(0, len(slice)-1, slice)
	// fmt.Println("sorted slice = ", slice)

	var slice []int = []int{8, 88, 555, 7, 91, 35, 66}
	quickSort(0, len(slice)-1, slice)
	fmt.Println("sorted slice = ", slice)

	// var slice []int = []int{555, 66, 91, 35, 88}
	// quickSort(0, len(slice)-1, slice)
	// fmt.Println("sorted slice = ", slice)

}
