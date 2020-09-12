package model

import "fmt"

// 声明 Customer 结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id int, name string,
	gender string, age int, phone string,
	email string) *Customer {

	customer := &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}

	return customer

}

// 返回用户的字符串信息
func (this *Customer) GetInfo() string {

	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		this.Id, this.Name, this.Gender, this.Age,
		this.Phone, this.Email)

	return info

}
