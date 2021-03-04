package main

import "fmt"

func main() {

	var a int = 300
	var b int = 400
	var ptr *int = &a
	*ptr = 100
	ptr = &b
	*ptr = 200

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*ptr)

	// a = 100
	// b = 200
	// *ptr = 200

}
