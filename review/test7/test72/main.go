/*
	复习
	switch 的穿透 fallthrought
*/
package main

import "fmt"

func main() {

	var num int = 10

	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到...")
	}

}
