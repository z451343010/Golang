/*
	复习
*/
package main

import "fmt"

func main() {

	var a int = 300
	temp := &a
	fmt.Printf("%T\n", temp)
	fmt.Printf("%T", &a)

}
