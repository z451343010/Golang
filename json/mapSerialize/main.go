package main

import (
	"encoding/json"
	"fmt"
)

// 序列化 map
func main() {

	// 定义一个map
	// key 为 string 类型
	// value 为 interface{} 类型，说明可以是任意类型
	var a map[string]interface{}

	// 使用map，需要 make
	a = make(map[string]interface{})
	a["name"] = "zjuen"
	a["age"] = 24
	a["address"] = "泉州"

	// 序列化 map
	// map 本身就是引用类型
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误 err = %v\n", err)
	}

	// 输出序列化后的结果
	fmt.Println(string(data))
}
