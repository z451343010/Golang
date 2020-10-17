/*
	Golang的反射入门
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
	Sex    string `json:"sex"`
}

func main() {

	monster := Monster{

		Name:   "Diablo",
		Age:    20,
		Salary: 200000,
		Sex:    "male",
	}

	data, _ := json.Marshal(monster)
	fmt.Println("json result = ", string(data))

}
