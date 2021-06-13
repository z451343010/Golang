/*
	牛客网
	华为机试
	HJ10
*/
package main

import "fmt"

func main() {

	var inputStr string
	fmt.Scanln(&inputStr)

	var byteSlice []byte

	for i := 0; i < len(inputStr); i++ {

		if inputStr[i] >= 0 && inputStr[i] <= 127 {

			if i == 0 {
				byteSlice = append(byteSlice, inputStr[i])
			} else {

				var flag bool

				for _, value := range byteSlice {

					if inputStr[i] == value {
						flag = true
						break
					}

				}

				if !flag {
					byteSlice = append(byteSlice, inputStr[i])
				}

			}

		}

	}

	fmt.Println(len(byteSlice))

}
