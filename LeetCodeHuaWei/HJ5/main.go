/*
	牛客网
	华为机试
	HJ5
	解题通过
*/
package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {

	var testStr string

	for {

		_, err := fmt.Scanln(&testStr)
		if err == io.EOF {
			break
		}

		var num int = 0
		for i := len(testStr) - 1; i > 1; i-- {

			var tempNum int
			if testStr[i] == 'A' {
				tempNum = 10
			} else {
				if testStr[i] > 'A' {
					tempNum = (int(testStr[i]) - int('A')) + 10
				} else {
					tempNum64, _ := strconv.ParseInt(string(testStr[i]), 10, 64)
					tempNum = int(tempNum64)
				}

			}

			index := len(testStr) - 1 - i
			if index == 0 {
				num += tempNum
			} else {

				for index > 0 {
					tempNum *= 16
					index--
				}
				num += tempNum

			}
		}
		fmt.Println(num)

	}

}
