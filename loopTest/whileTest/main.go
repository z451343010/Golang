package main

import "fmt"

// Go语言没有while和do...while语法
// 如果我们需要使用类似于其它语言
// （如：Java/c 的while和do...while），
// 则可以通过for循环来实现其使用效果。
func main() {

	var i byte = 0

	for {

		if i >= 10 {
			break
		}
		fmt.Println("Hello World!")

	}

}
