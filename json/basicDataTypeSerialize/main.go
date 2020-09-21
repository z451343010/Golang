package main

import (
	"encoding/json"
	"fmt"
)

// 对基本数据类型进行序列化
func main() {

	var num float64 = 123.45699

	data, err := json.Marshal(num)
	if err != nil {
		fmt.Println("serialize err = ", err)
	}

	fmt.Println(string(data))

}
