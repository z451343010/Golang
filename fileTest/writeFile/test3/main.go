package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	打开一个存在的文件
	在原来的内容追加 "ABC!ENGLISH"
*/

func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test1\\testFileWrite.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}

	// 准备写入
	writeStr := "ABC!ENGLISH\n"

	writer := bufio.NewWriter(file)
	writer.WriteString(writeStr)

	writer.Flush()

	defer file.Close()

}
