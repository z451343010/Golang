/*
	选择排序
	思路：循环取最小值，然后把最小值以此放入数组前部
*/

package main

import "fmt"

func selectionSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {

		var min int = slice[i]
		var minIndex int = i

		for j := i + 1; j < len(slice); j++ {

			if slice[j] < min {
				min = slice[j]
				minIndex = j
			}

		}

		slice[minIndex], slice[i] = slice[i], slice[minIndex]

	}

}

func main() {

	var slice []int = []int{108, 3, 88, 12, 8, 88, 555, 7, 91, 35, 66}
	selectionSort(slice)
	fmt.Println("slice = ", slice)

}
