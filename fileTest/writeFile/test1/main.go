package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	用带缓存的方式写入文件
*/
func main() {

	// 创建一个文件，写入内容5句 "Hello Zhang"
	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test1\\testFileWrite.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}

	// 准备写入
	writeStr := "Hello Zhang\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// 把文件的内容先写入缓存
		writer.WriteString(writeStr)
	}

	// 因为 writer 是带缓存的，因此在调用 WriterString()时
	// 其实内容是先写到缓存的，所以需要调用 Flush()
	// 将缓存的数据真正的写入到文件中
	// 否则文件中会没有数据
	writer.Flush()

	defer file.Close()

}
