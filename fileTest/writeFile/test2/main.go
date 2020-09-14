package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	打开一个存在的文件
	将原来的内容覆盖成新的10句
	"Hello Xanxuskkk"
*/

func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test1\\testFileWrite.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}

	// 准备写入
	writeStr := "Hello Xanxuskkk\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		// 把文件的内容先写入缓存
		writer.WriteString(writeStr)
	}

	writer.Flush()

	defer file.Close()

}
