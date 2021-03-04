/*
	复习
	switch 语句
*/
package main

import "fmt"

func main() {

	switch 1 + 2 {
	case 3:
		fmt.Println(3)
	}

	var testNum int = 4
	switch testNum {
	case testNum:
		fmt.Println(5)
	case 4:
		fmt.Println(4)
	}

}
