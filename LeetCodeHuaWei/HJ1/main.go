package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var testStr string
	reader := bufio.NewReader(os.Stdin)
	testStr, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count int = 0
	for i := len(testStr) - 2; i >= 0; i-- {

		if testStr[i] == ' ' {
			break
		}
		count++

	}
	fmt.Println(count - 1)

}
