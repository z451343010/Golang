package main

import (
	"encoding/json"
	"fmt"
)

/*
	json反序列化结构体
*/

// 定义一个结构体
type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

func main() {

	str := "{\"name\":\"Goku\",\"age\":24,\"birthday\":" +
		"\"1995-11-07\",\"salary\":1000000,\"skill\":\"Programing\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unserialize err = ", err)
	}

	fmt.Println(monster)

}
