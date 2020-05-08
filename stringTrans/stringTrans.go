package main

import (
	"fmt"
	"strconv"
)

// string 转换成基本数据类型
func main() {

	var str string = "true"
	var flag bool

	// 1.函数 strconv.ParseBool() 会返回两个值(value bool, err error)
	// 2.我们只想获取到 value bool，不想获取 err
	// 因此使用_来忽略
	flag, _ = strconv.ParseBool(str)
	fmt.Printf("flag type is %T and value is %t \n", flag, flag)

	var str2 string = "1234"
	var n1 int64

	// 字符串，进制，64位
	// 0：int类型 8：int8 16：int16 64：int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type is %T and n1 value is %v \n", n1, n1)

	// float类型
	var floatNumStr string = "123456.123456"
	var floatNum float64
	floatNum, _ = strconv.ParseFloat(floatNumStr, 64)
	fmt.Printf("floatNum type is %T and floatNum value is %v \n", floatNum, floatNum)

	// 将一个非数字的字符串转换成整型
	// 如：把字符串 "hello" 转成10进制64位的整数
	var strHello string = "hello"
	var intHello int64
	intHello, _ = strconv.ParseInt(strHello, 10, 64)
	fmt.Printf("intHello type %T value = %v \n", intHello, intHello)

	// 把一个bool类型的变量赋值后，再转换，看看最后输出的结果
	var transFlag bool = true
	var transStr = "hello"
	// fmt.Println(transFlag)
	transFlag, _ = strconv.ParseBool(transStr)
	fmt.Printf("transFlag type is %T and value is %v \n", transFlag, transFlag)
	// fmt.Println(transFlag)

	// 获取变量的地址
	var num int = 3
	fmt.Println("the address of num is ", &num)

}
