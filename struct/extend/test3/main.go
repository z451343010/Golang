package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

func (b *B) SayOk() {
	fmt.Println("B SayOk")
}

type B struct {
	A
}

func main() {
	var b B
	// fmt.Println(b.A.Name)
	// fmt.Println(b.A.age)
	// b.A.SayOk()
	// b.A.hello()

	// fmt.Println(b.Name)
	// fmt.Println(b.age)

	// 通过匿名结构体来区分
	b.A.SayOk()

}
