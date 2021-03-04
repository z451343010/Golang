package main

import "fmt"

func test(char byte) byte {
	return char + 1
}

func main() {

	var theChar byte

	fmt.Println("请输入字符：")
	fmt.Scanf("%c", &theChar)
	// theChar = 'a'

	switch test(theChar) {
	case 'a':
		fmt.Println("Monday")
	case 'b':
		fmt.Println("Tuesday")
	}

}
