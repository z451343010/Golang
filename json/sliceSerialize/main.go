package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "zhang"
	m1["age"] = 24
	m1["address"] = "Quanzhou"

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "zeng"
	m2["age"] = 24
	m2["address"] = "Liuzhou"

	slice = append(slice, m1)
	slice = append(slice, m2)

	// 序列化 slice
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误 err = %v\n", err)
	}

	// 输出序列化后的结果
	fmt.Println(string(data))

}
