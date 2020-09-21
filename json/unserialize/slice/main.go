package main

import (
	"encoding/json"
	"fmt"
)

// 反序列化 slice
func main() {

	str := "[{\"address\":\"Beijing\",\"age\":\"24\"," +
		"\"name\":\"zhang\"},{\"address\":[\"NewYork\",\"Paris\"],\"age\":\"24\",\"name\":\"zeng\"}]"

	// 反序列化 slice
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("unmarshal err = ", err)
	}

	fmt.Println(slice)

}
