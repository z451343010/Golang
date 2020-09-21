package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

func main() {

	monster := Monster{
		Name:     "Goku",
		Age:      24,
		Birthday: "1995-11-07",
		Salary:   1000000,
		Skill:    "Programing",
	}

	// 序列化 monster
	data, err := json.Marshal(&monster)

	if err != nil {
		fmt.Println("serialize err = ", err)
	}

	fmt.Println(string(data))

}
