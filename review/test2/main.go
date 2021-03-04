/*
	复习
	可以用单个字节来表示汉字
*/
package main

import "fmt"

func main() {

	var chineseCharacter int = '北'
	fmt.Println(chineseCharacter)
	fmt.Printf("%c\n", chineseCharacter)

	var testC int = 1000
	fmt.Printf("%c\n", testC)

}
