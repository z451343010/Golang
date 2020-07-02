package main

import "fmt"

// 定义一个 Cat结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {

	var cat1 Cat
	cat1.Name = "琦琦"
	cat1.Age = 3
	cat1.Color = "虎斑色"

	fmt.Println(cat1)

}
