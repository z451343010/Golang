package main

import (
	"fmt"
	"io/ioutil"
)

/*
	编写一个程序，将一个文件的内容，写入到另一个文件。
	注释：这两个文件已经存在了
	说明：使用 ioutil.ReadFile 和 iotuil.WriteFile
*/
func main() {

	filePath1 := "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test6\\file1.txt"

	filePath2 := "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test6\\file2.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		// 说明读取文件有错误
		fmt.Println("read file err = ", err)
		return
	}

	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
	}

}
