package main

import (
	"fmt"
	"io/ioutil"
)

/*
	一次性读取文件中的所有内容
	不带缓冲
*/
func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src\\fileTest\\test1\\test1.txt"

	// 使用ioutil.ReadFile() 一次性将文件读取到位
	// ReadFile() 函数里面就包含了文件的打开和关闭
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Printf("read file err = %v ", err)
	}

	// 把读取到的内容显示到终端
	fmt.Printf("%v", content) // []byte

	// 以string的形式输出具体内容
	fmt.Printf("%v", string(content))

}
