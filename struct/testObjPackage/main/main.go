package main

import (
	"fmt"
	"struct/testObjPackage/model"
)

func main() {

	person := model.NewPerson("zjune")
	fmt.Println(person)

	(*person).SetAge(15)
	fmt.Println(person)

	var age int
	age = (*person).GetAge()
	fmt.Println(age)

}
