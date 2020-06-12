package main

import "fmt"

// 冒泡排序
func main() {

	var numArr [6]int = [6]int{24, 99, 80, 57, 13, 69}

	for i := len(numArr) - 1; i >= 0; i-- {

		for j := 0; j < i; j++ {

			if numArr[j] > numArr[j+1] {

				var temp int
				temp = numArr[j]
				numArr[j] = numArr[j+1]
				numArr[j+1] = temp

			}

		}

	}

	fmt.Println(numArr)

}
