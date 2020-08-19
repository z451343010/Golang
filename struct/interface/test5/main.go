package main

import "fmt"

type Stu struct {
}

type BInterface interface {
	test1()
}

type CInterface interface {
	test2()
}

type AInterface interface {
	BInterface
	CInterface
	test3()
}

func (stu Stu) test1() {
	fmt.Println("stu test1")
}

func (stu Stu) test2() {
	fmt.Println("stu test2")
}

func (stu Stu) test3() {
	fmt.Println("stu test3")
}

func main() {

	var stu Stu
	var a AInterface = stu
	a.test1()
	// a.test2()
	// a.test3()

	var b BInterface = stu
	b.test1()

}
