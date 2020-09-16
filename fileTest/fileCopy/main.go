package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	文件拷贝
*/

func CopyFile(srcFileName string, dstFileName string) (written int64, err error) {

	srcFile, srcErr := os.Open(srcFileName)

	if err != nil {
		fmt.Println("open file srcErr = ", srcErr)
		return
	}

	defer srcFile.Close()

	// 通过srcFile，获取 Reader
	reader := bufio.NewReader(srcFile)

	// 打开 dstFileName
	dstFile, dstErr := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if dstErr != nil {
		fmt.Println("open file dstErr = ", dstErr)
		return
	}

	// 通过dstFile,获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	// 调用 io 包内置的 Copy() 函数实现复制
	return io.Copy(writer, reader)

}

func main() {

	// 源文件路径
	srcFileName := "F:\\Golang\\goProgramProject\\src\\" +
		"fileTest\\fileCopy\\Einstein.jpg"
	// 目标路径
	dstFileName := "F:\\Golang\\goProgramProject\\src\\" +
		"fileTest\\fileCopy\\picture\\relativity.jpg"

	_, error := CopyFile(srcFileName, dstFileName)

	if error != nil {
		fmt.Println("copy file error = ", error)
	}

}
