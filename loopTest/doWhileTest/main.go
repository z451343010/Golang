package main

import "fmt"

// 使用 do...while 的实现完成输出10句 "Hello World"
func main() {

	var i int = 0

	for {

		fmt.Println("Hello World")
		i++
		if i >= 10 {
			break
		}

	}

}
