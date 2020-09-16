package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int // 统计英文的个数
	NumCount   int // 统计数字的个数
	SpaceCount int // 统计空格的个数
	OtherCount int // 统计其它字符的个数
}

/*
	统计英文、数字、空格和其它字符数量
	说明：统计一个文件中含有的英文、数字、空格以及其它字符数量
*/
func main() {

	fileName := "F:\\Golang\\goProgramProject\\src\\" +
		"fileTest\\fileExc\\test1\\count.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var charCount CharCount

	for {

		str, err := reader.ReadString('\n')

		for _, value := range str {

			switch {
			case value >= 'a' && value <= 'z':
				charCount.ChCount++
			case value >= 'A' && value <= 'Z':
				charCount.ChCount++
			case value >= '0' && value <= '9':
				charCount.NumCount++
			case value == '\n' || value == '\t':
				charCount.SpaceCount++
			default:
				charCount.OtherCount++

			}

		}

		if err == io.EOF {
			break
		}

	}

	fmt.Println("该文件的英文字母数量为：", charCount.ChCount)
	fmt.Println("该文件的数字数量为：", charCount.NumCount)
	fmt.Println("该文件的空格数量为：", charCount.SpaceCount)
	fmt.Println("该文件的其它字符数量为：", charCount.OtherCount)

}
