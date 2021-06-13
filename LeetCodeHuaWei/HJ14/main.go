/*
	牛客网
	华为机试
	HJ14
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	var wordNum int
	fmt.Scanln(&wordNum)

	var wordSlice []string

	for i := 0; i < wordNum; i++ {

		var word string
		fmt.Scanln(&word)
		wordSlice = append(wordSlice, word)

	}

	sort.Sort(sort.StringSlice(wordSlice))

	for _, value := range wordSlice {
		fmt.Println(value)
	}

}
