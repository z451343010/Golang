package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int
	n int
}

func main() {

	var e E
	e.Name = "zjune"
	e.Age = 200
	e.int = 20
	e.n = 40
	fmt.Println(e)

}
