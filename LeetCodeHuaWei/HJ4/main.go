/*
	牛客网
	华为机试
	HJ4
	实际输出和预期输出一致
	但是提交的时候报“格式不正确”
*/
package main

import (
	"fmt"
	"io"
)

func divideStr(str string) {

	if len(str) <= 8 {
		var strArr [8]string
		for i := 0; i < len(strArr); i++ {
			if i < len(str) {
				strArr[i] = string(str[i])
			} else {
				strArr[i] = "0"
			}
		}

		for i := 0; i < len(strArr); i++ {
			fmt.Print(strArr[i])
		}
		//fmt.Println("len = ", len(strArr))

	} else {

		divide := len(str) / 8
		remain := len(str) % 8

		var strArr []string
		if remain != 0 {
			strArr = make([]string, (divide+1)*8)
		} else {
			strArr = make([]string, divide*8)
		}

		for i := 0; i < divide; i++ {
			tempI := i + 1
			for j := i * 8; j < tempI*8; j++ {
				strArr[j] = string(str[j])
			}
		}
		for i := divide * 8; i < len(strArr); i++ {
			if i < len(str) {
				strArr[i] = string(str[i])
			} else {
				strArr[i] = "0"
			}
		}
		for i := 0; i < len(strArr); i++ {
			if i%8 == 0 {
				fmt.Println()
			}
			fmt.Print(strArr[i])
		}

	}

}

func main() {

	var testStr string
	for {
		_, err := fmt.Scanln(&testStr)
		if err == io.EOF {
			break
		}
		divideStr(testStr)
	}

}
