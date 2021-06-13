package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	testStr, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tempStr := strings.ToUpper(string(testStr))

	var testChar byte
	fmt.Scanf("%c", &testChar)

	count := 0
	for i := 0; i < len(tempStr)-1; i++ {

		if string(tempStr[i]) == strings.ToUpper(string(testChar)) {
			count++
		}

	}

	fmt.Println(count)

}
