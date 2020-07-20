package main

import "fmt"

/*
	Golang在创建结构体变量时，可以直接指定字段的值
*/

type Student struct {
	Name string
	Age  int
}

func main() {

	var stu Student = Student{
		Name: "marry",
		Age:  30,
	}

	stu2 := Student{
		Name: "marry",
		Age:  30,
	}

	var stu3 *Student = &Student{"smith", 30}

	var stu4 *Student = &Student{
		Name: "Scott",
		Age:  80,
	}

	fmt.Println(stu)
	fmt.Println(stu2)
	fmt.Println(*(stu3))
	fmt.Println(*(stu4))

}
