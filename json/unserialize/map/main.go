package main

import (
	"encoding/json"
	"fmt"
)

// map 反序列化
func main() {

	str := "{\"name\":\"zhang\",\"age\":\"30\",\"address\":\"Hangzhou\"}"

	// 定义一个map
	var testMap map[string]interface{}

	// JSON反序列化成map
	// 反序列化操作时不需要make
	// 因为make操作已经被封装到了Unmarshal()
	err := json.Unmarshal([]byte(str), &testMap)
	if err != nil {
		fmt.Println("unmarshal err = ", err)
	}

	fmt.Println(testMap)

}
