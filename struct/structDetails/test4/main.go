package main

import (
	"encoding/json"
	"fmt"
)

// struct的每个字段上，可以写上一个tag
// 该tag可以通过反射机制获取
// 常见的使用场景就是序列化和反序列化。

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	var monster Monster
	monster.Name = "red man"
	monster.Age = 10
	monster.Skill = "fire"

	// 我们希望将其进行json格式的序列化处理
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json encoding err ", err)
		return
	}

	fmt.Println("json后的数据", string(data))

}
