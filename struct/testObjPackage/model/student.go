package model

import "fmt"

type person struct {
	Name string
	// 其它包不能直接访问
	age    int
	salary float64
}

// 写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问 age，编写一对Set和Get方法
func (person *person) SetAge(age int) {

	if age > 0 && age < 150 {
		(*person).age = age
	} else {
		fmt.Println("输入的年龄范围不正确")
	}

}

func (person *person) GetAge() int {
	return (*person).age
}
