package main

import "fmt"

// map的使用方式
// 演示一个 key-value 的value是map的案例
// 比如：我们要存放3个学生的信息
// 每个学生都有name和sex信息
func main() {

	stuInfo := make(map[string]map[string]string)
	stuInfo["no1"] = make(map[string]string, 2)
	stuInfo["no1"]["name"] = "tom"
	stuInfo["no1"]["sex"] = "male"

	fmt.Println(stuInfo)

}
