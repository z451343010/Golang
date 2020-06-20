package main

import "fmt"

// map切片
func main() {

	// 声明一个map切片
	var heros []map[string]string
	heros = make([]map[string]string, 2)

	if heros[0] == nil {
		heros[0] = make(map[string]string, 2)
		heros[0]["name"] = "Nevermore"
		heros[0]["speed"] = "400"
	}

	if heros[1] == nil {
		heros[1] = make(map[string]string, 2)
		heros[1]["name"] = "Zeus"
		heros[1]["speed"] = "300"
	}

	// 这里我们需要使用到切片的append函数，来动态增加
	newHero := map[string]string{
		"name":  "roshan",
		"speed": "100",
	}

	heros = append(heros, newHero)

	fmt.Println(heros)

}
