package main

import "fmt"

// 使用switch把小写类型的char型转换为大写（键盘输入）
// 只转换 a,b,c,d,e
// 其它的输出"other"
func main() {

	var letter byte
	fmt.Println("请输入需要转换的字母：")
	fmt.Scanf("%c", &letter)

	switch letter {
	case 'a', 'b', 'c', 'd', 'e':
		letter -= 32
		fmt.Printf("%c", letter)
	default:
		fmt.Println("other")
	}

}
