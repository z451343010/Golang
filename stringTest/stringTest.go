package main

func main() {
	var address string = "Beijing Greatwall"
	println(address)
	address = "qinghua"
	println(address)

	// var str string = "zjune"
	// str[0] = 'c '

	// 字符串的两种表现形式：
	// 1.双引号，会识别转义字符
	// 2.反引号，以字符串的原生形式输出，包括换行和特殊字符
	// 可以实现防止攻击、输出源代码等效果
	var str2 string = `
	var address string = "Beijing Greatwall"
	println(address)
	address = "qinghua"
	println(address)	
	`
	println(str2)

}
