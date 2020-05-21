package main

import "fmt"

// 演示goto的使用
func main() {

	fmt.Println("ok1")
	goto label
	fmt.Println("ok2")
	fmt.Println("ok3")
label:
	fmt.Println("ok4")
	fmt.Println("ok5")

}
