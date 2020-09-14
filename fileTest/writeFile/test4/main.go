package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	打开一个存在的文件
	将原来的内容读出显示在终端，并且追加5句
	"hello,北京"
*/

func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test1\\testFileWrite.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}

	// 循环读取文件的内容

	reader := bufio.NewReader(file)

	for {

		str, err := reader.ReadString('\n') // 读到换行符就结束
		if err == io.EOF {                  // io.EOF 表示文件末尾,EOF是 package io 的常量
			break
		}

		// 输出内容
		fmt.Print(str) // 本身就会读取文件中的换行符

	}

	// 准备写入
	writeStr := "hello,北京\n"

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(writeStr)
	}

	writer.Flush()

	defer file.Close()

}
