package main

import "fmt"

// for循环的第三种使用方式
func main() {

	var i int = 0
	for {

		fmt.Println("hello")
		i++
		if i > 9 {
			break
		}
	}

}
