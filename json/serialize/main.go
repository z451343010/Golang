package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
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
