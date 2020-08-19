package main

import "fmt"

type AInterface interface {
	Say()
}

type Student struct {
	Name string
}

func (stu *Student) Say() {
	fmt.Println("stu say()")
}

func main() {
	var stu Student
	var a AInterface = &stu
	a.Say()
}
