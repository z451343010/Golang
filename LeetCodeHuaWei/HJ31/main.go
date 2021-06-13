/*
	牛客网
	华为机试
	HJ31
	单词倒排（困难）
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}

	var newStr string
	for i := 0; i < len(line)-2; i++ {
		if (line[i] >= 97 && line[i] <= 122) || (line[i] >= 65 &&
			line[i] <= 90) {
			newStr += string(line[i])
		} else {
			newStr += string(" ")
		}
	}

	strArr := strings.Split(newStr, " ")
	for i := len(strArr) - 1; i >= 0; i-- {
		if strArr[i] != " " {
			fmt.Print(strArr[i], " ")
		}
	}

}
