/*
	复习
*/
package main

import "fmt"

func test02(n1 int) {

	n1 = n1 + 10
	fmt.Println("test02() n1 = ", n1)

}

func main() {

	n1 := 20
	test02(n1)
	fmt.Println("main() n1 = ", n1)

}
