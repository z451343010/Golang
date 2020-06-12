package main

import "fmt"

// string和slice
func main() {

	var str string = "hello@zjune"
	slice := str[6:]
	fmt.Println(slice)

	// 利用切片
	// 将字符串 hello@zjune 改为 hello@xjune
	slice2 := []byte(str)
	slice2[6] = 'x'
	str = string(slice2)
	fmt.Println(str)

	// 我们转成[]byte后，可以处理英文和数字，但是不能处理中文。
	// 原因是[]byte用字节来处理，而一个汉字，是3个字节，因此就会出现乱码
	// 解决方法：将string转成rune[]即可。因为rune[]是按字符来计算的。
	slice3 := []rune(str)
	slice3[6] = '张'
	str = string(slice3)
	fmt.Println(str)

}
