package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src\\fileTest\\test1\\test1.txt"

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("open file err = ", err)
	}

	// 函数退出时及时的关闭文件
	// 需要及时关闭 file 句柄，否则会有内存泄漏
	// defer 关键字在函数退出的时候执行
	defer file.Close()

	// 创建一个带缓冲的 *Reader
	/*
		const {
			defaultBufSize = 4096  // 默认的缓冲区为4096
		}
	*/

	reader := bufio.NewReader(file)

	// 循环读取文件的内容
	for {

		str, err := reader.ReadString('\n') // 读到换行符就结束
		if err == io.EOF {                  // io.EOF 表示文件末尾,EOF是 package io 的常量
			break
		}

		// 输出内容
		fmt.Print(str) // 本身就会读取文件中的换行符

	}

}
