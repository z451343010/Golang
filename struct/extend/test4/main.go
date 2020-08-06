package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name  string
	score int
}

type C struct {
	A
	B
}

type D struct {
	a A
}

func main() {

	var c C
	c.A.Name = "kkk"

	fmt.Println(c.A.Name)

	var d D
	d.a.Name = "jack"

}
