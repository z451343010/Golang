/*
	牛客网
	华为机试
	HJ12
*/
package main

import "fmt"

func main() {

	var inputStr string
	fmt.Scanln(&inputStr)

	for i := len(inputStr) - 1; i >= 0; i-- {
		fmt.Printf("%c", inputStr[i])
	}

}
