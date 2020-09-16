package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	编写一个程序，将一个文件的内容，写入到另一个文件。
	注释：这两个文件已经存在了
*/
func main() {

	filePath1 := "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test5\\file1.txt"

	filePath2 := "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test5\\file2.txt"

	file1, err1 := os.OpenFile(filePath1, os.O_RDWR, 0666)
	file2, err2 := os.OpenFile(filePath2, os.O_RDWR|os.O_TRUNC, 0666)

	if err1 != nil {
		fmt.Println("open file err = ", err1)
	}

	if err2 != nil {
		fmt.Println("open file err = ", err2)
	}

	reader := bufio.NewReader(file1)
	writer := bufio.NewWriter(file2)

	// 循环读取文件的内容
	for {

		strRd, errRd := reader.ReadString('\n') // 读到换行符就结束

		fmt.Print(strRd)

		// 写入缓冲区
		writer.WriteString(strRd)

		if errRd == io.EOF { // io.EOF 表示文件末尾,EOF是 package io 的常量
			break
		}

	}

	writer.Flush()

	defer file1.Close()
	defer file2.Close()

}
