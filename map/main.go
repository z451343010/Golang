package main

import "fmt"

// map的入门案例
func main() {

	// 声明
	var a map[string]string

	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	// 10说明这个map最大可以分配10个键值对
	a = make(map[string]string, 10)
	a["NO1"] = "zzz"
	a["NO2"] = "xxx"
	fmt.Println(a)

}
