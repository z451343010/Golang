/*
	牛客网
	华为机试
	HJ13
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	var begin int
	var strSlice []string

	for i := 0; i <= len(line)-2; i++ {

		var word string

		if i == 0 {
			begin = 0
		}

		if line[i] == 32 || (i == len(line)-2) {

			for j := begin; j < i; j++ {
				//fmt.Println(line[j])
				word += fmt.Sprintf("%c", line[j])
			}

			begin = i + 1

			strSlice = append(strSlice, word)

		}
		// fmt.Println(line[i])

	}

	for i := len(strSlice) - 1; i >= 0; i-- {
		fmt.Printf("%s ", strSlice[i])
	}

}
