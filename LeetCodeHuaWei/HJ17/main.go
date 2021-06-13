/*
	牛客网
	华为机试
	HJ17
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	fmt.Scanln(&inputStr)

	strArr := strings.Split(inputStr, ";")

	var indexI, indexJ int
	for i := 0; i < len(strArr)-1; i++ {
		// fmt.Println(strArr[i])

		if len(strArr[i]) >= 2 {

			tempChar := strArr[i][:1]
			tempStr := strArr[i][1:]
			tempNum, err := strconv.Atoi(tempStr)

			if err == nil {
				switch tempChar {
				case "A":
					indexI -= tempNum
				case "D":
					indexI += tempNum
				case "W":
					indexJ += tempNum
				case "S":
					indexJ -= tempNum
				}
			}

		}

	}

	fmt.Printf("%d,%d", indexI, indexJ)

}
