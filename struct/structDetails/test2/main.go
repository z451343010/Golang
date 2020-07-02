package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {

	var a A
	var b B

	a = A(b)
	fmt.Println(a, b)

}
