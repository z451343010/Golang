package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) speak() {

	person.Name = "jack"
	fmt.Println(person.Name, "is a good man")

}

func main() {

	var person Person
	person.speak()

}
