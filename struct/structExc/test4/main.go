package main

import "fmt"

type Person struct {
}

func (person Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {

	n1 := 10
	n2 := 20
	var person Person
	result := person.getSum(n1, n2)
	fmt.Println(result)

}
