package main

import (
	"fmt"
)

func main() {
	var randNum int
	fmt.Scan(&randNum)
	var numSlice []int
	for i := 0; i < randNum; i++ {
		var num int
		fmt.Scan(&num)
		numSlice = append(numSlice, num)
	}
	fmt.Println(numSlice)
	// for i := 0; i < len(numSlice)-1; i++ {
	// 	for k := 0; i < len(numSlice)-1-i; k++ {
	// 		if numSlice[i] > numSlice[i+1] {
	// 			numSlice[i], numSlice[i+1] = numSlice[i+1], numSlice[i]
	// 		}

	// 	}
	// }
	fmt.Println(numSlice)

}
