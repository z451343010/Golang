package main

import "fmt"

/*
	对于方法（如struct的方法），接收者为值类型时
	可以直接用指针类型的变量调用方法，反过来同样也可以。
*/

type Person struct {
	Name string
}

func (person Person) test3() {

	person.Name = "jack"
	fmt.Println("test3() = ", person.Name)

}

func main() {

	var person Person
	person.Name = "rose"

	// 形式上看是传入地址，本质上还是值拷贝
	(&person).test3()

	fmt.Println("main() Name =", person.Name)

}
