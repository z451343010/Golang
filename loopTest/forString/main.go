package main

import "fmt"

// 用for循环遍历字符串
// len() 求长度
func main() {

	var testStr string = "Hello,World!"
	for i := 0; i < len(testStr); i++ {
		fmt.Printf("%c\n", testStr[i])
	}

}
