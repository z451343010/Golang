package main

import (
	"fmt"
	"os"
)

/*
	判断文件或者目录是否存在
*/
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, nil

}

func main() {

	var filePath string
	filePath = "F:\\Golang\\goProgramProject\\src" +
		"\\fileTest\\writeFile\\test6"

	flag, err := PathExists(filePath)

	fmt.Println(flag)
	fmt.Println(err)

}
