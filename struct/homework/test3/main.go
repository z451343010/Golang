package main

import "fmt"

/*
	景区门票案例
*/
type Visitor struct {
	name        string
	age         int
	ticketPrice float64
}

func (visitor Visitor) getTicketPrice() float64 {

	if visitor.age >= 18 {
		visitor.ticketPrice = 0
	} else {
		visitor.ticketPrice = 20
	}

	return visitor.ticketPrice

}

func main() {

	var visitor Visitor
	var name string
	var age int
	var ticketPrice float64

	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	ticketPrice = visitor.getTicketPrice()
	fmt.Println(name, "的年龄为:", age, "票价为：", ticketPrice)

}
