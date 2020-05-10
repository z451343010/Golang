package main

import "fmt"

// 课堂练习2
// 定义一个变量保存华氏度，求出华氏温度对应的摄氏温度
func main() {

	var fahrenheit float32 = 134.2
	var celsius float32
	celsius = (fahrenheit - 100) * 5.0 / 9
	fmt.Println(celsius)

}
