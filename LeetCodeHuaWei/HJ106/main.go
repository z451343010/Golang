/*
	牛客网
	华为机试
	HJ106
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := len(line) - 2; i >= 0; i-- {
		fmt.Printf("%c", line[i])
	}

}
