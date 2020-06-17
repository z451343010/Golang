package main

import "fmt"

// map使用的几种方式
func main() {

	// 方式一
	var a map[string]string
	a = make(map[string]string, 10)
	a["NO1"] = "zzz"
	a["NO2"] = "xxx"
	fmt.Println(a)

	// 方式二
	cities := make(map[string]string)
	cities["NO1"] = "北京"
	cities["NO2"] = "上海"
	cities["NO3"] = "杭州"

	fmt.Println(cities)

	// 方式三
	var heroes map[string]string = map[string]string{
		"hero1": "影魔",
		"hero2": "宙斯",
	}

	// 或者直接用类型推导
	heroes2 := map[string]string{
		"hero1": "影魔",
		"hero2": "宙斯",
	}

	fmt.Println(heroes)
	fmt.Println(heroes2)

}
