package main

import "fmt"

// 字符串遍历
// 同时处理有中文的问题
func main() {

	var str string = "hello北京"
	var tempStr = []rune(str)

	for i := 0; i < len(tempStr); i++ {
		fmt.Printf("%c\n", tempStr[i])
	}

}
