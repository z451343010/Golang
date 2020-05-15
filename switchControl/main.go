package main

import "fmt"

func test(char byte) byte {
	return char + 1
}

func main() {

	var theChar byte

	fmt.Println("请输入字符：")
	fmt.Scanf("%c", &theChar) // 用指定的格式接收键盘输入的字符

	switch test(theChar) {
	case 'a':
		fmt.Println("Monday")
	case 'b':
		fmt.Println("Tuesday")
	case 'c':
		fmt.Println("Wednesday")
	case 'd':
		fmt.Println("Thursday")
	case 'e':
		fmt.Println("Friday")
	case 'f':
		fmt.Println("Saturday")
	case 'g':
		fmt.Println("Sunday")
	default:
		fmt.Println("wrong day")
	}

}
