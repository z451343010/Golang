package main

import (
	"fmt"
	"os"
)

/*
	文件demo
*/

// F:\Golang\goProgramProject\src\fileTest\test1\test1.txt
func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src\\fileTest\\test1\\test1.txt"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	// 输出文件
	fmt.Printf("%v", file)

	err = file.Close()

	if err != nil {
		fmt.Println("close file err = ", err)
	}

}
