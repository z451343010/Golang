package main

import "fmt"

// 类名和字段名首字母大写可以实现跨包调用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {

	var p1 Person
	fmt.Println(p1)

	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	p1.map1 = make(map[string]string)
	p1.map1["NO1"] = "zjune"

	fmt.Println(p1)

	var p3 *Person = new(Person)
	(*p3).Name = "Dante"
	(*p3).Age = 200
	fmt.Println(*p3)

	var person *Person = &Person{}
	(*person).Name = "Viggle"
	(*person).Age = 200
	fmt.Println(*person)

}
