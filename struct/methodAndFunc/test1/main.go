package main

import "fmt"

/*
	函数
	对于普通函数，接收者为值类型时
	不能将指针类型的数据直接传递
	反之亦然
*/

type Person struct {
	Name string
}

func test1(person Person) {

	fmt.Println(person.Name)

}

func test2(person *Person) {

	fmt.Println(person.Name)

}

func main() {

	var person Person
	test1(person)

}
