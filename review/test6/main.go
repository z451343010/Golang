/*
	复习
	多分支的使用陷阱
*/
package main

import "fmt"

func main() {

	var num int8 = 10

	if num > 9 {
		fmt.Println("ok1")
	} else if num > 6 {
		fmt.Println("ok2")
	} else if num > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}

}
