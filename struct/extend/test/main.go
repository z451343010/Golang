package main

/*
	继承
	Golang的匿名结构体
*/

type Goods struct {
	Name  string
	Price float64
}

type Book struct {
	Goods
	Writer string
}

func main() {

}
