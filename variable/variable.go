package main

import (
	"fmt"
	"unsafe"
)

var v1 int // 全局变量

func main() {
	// var i int 声明变量
	// i = 10    赋值变量
	// fmt.Println(i) 使用变量

	// Golang变量使用的三种方式
	var v1 int  // 指定变量类型，声明后若不复制，使用默认值
	var v2 = 10 // 根据值自行判断变量类型（类型推导）
	v3 := 20    // 省略var

	var flag bool

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println("v3 = ", v3)
	// fmt.Println("global v1 = " +  v1)
	fmt.Println(flag)
	var n1 = 100
	fmt.Printf("n1 的类型 %T\n", n1)

	// 如何查看一个变量占用的字节大小和数据类型？
	var n2 int64 = 10

	// unsafe.Sizeof() 可以返回变量占用的字节数
	fmt.Printf("n2 的 类型 %T n2占用的字节数是 %d", n1, unsafe.Sizeof(n2))

	// 小数类型
	var price float32 = 10.25
	fmt.Println(price)

	// 浮点数尾数部分可能造成精度损失
	var f1 float32 = -123.0000901
	var f2 float64 = -123.0000901
	fmt.Println(f1)
	fmt.Println(f2)

	// Golang字符
	var testChar byte = 'c'
	fmt.Println("testChar = ", testChar)

	// 如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("testChar = %c \n", testChar)

	// 可以用int类型来表示单个汉字
	var chineseCharacter int = '北'
	fmt.Printf("%c \n", chineseCharacter)

	// 给某个变量赋一个整数值
	// 格式化输出
	// 这个整数在UTF-8编码中对应的字符
	var testC int = 500
	fmt.Printf("%c \n", testC)

}
