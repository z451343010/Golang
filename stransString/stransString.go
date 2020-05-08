package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型和string的转换
func main() {

	var num1 int = 88
	var num2 float64 = 88.88
	var b bool
	var myChar byte = 'h'

	// 方式一，使用 fmt.Sprintf()
	var str string = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q \n", str, str)

	var str2 string = fmt.Sprintf("%f", num2)
	fmt.Printf("str2 type %T str2 = %q \n", str2, str2)

	var str3 string = fmt.Sprintf("%t", b)
	fmt.Printf("str3 type %T str3 = %q \n", str3, str3)

	var str4 string = fmt.Sprintf("%c", myChar)
	fmt.Printf("str4 type %T str4 = %q \n", str4, str4)

	// 方式二，使用 strconv 函数
	fmt.Println()
	fmt.Println("use strconv function ---------")

	var str5 string = strconv.FormatInt(int64(num1), 10) // 第二个参数表示进制
	fmt.Printf("str5 type %T str5 = %q \n", str5, str5)

	// f代表格式（-ddd.dddd）,第三个参数代表精度，表示小数点保留10位
	// 64 表示小数类型是 float64（双精度）
	var str6 string = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str6 type %T str6 = %q \n", str6, str6)

	var str7 string = strconv.FormatBool(b)
	fmt.Printf("str7 type %T str7 = %q \n", str7, str7)

	// strconv包中有一个函数Itoa
	var num3 int = 3379
	var str8 string = strconv.Itoa(num3)
	fmt.Printf("str8 type %T str8 = %q \n", str8, str8)

}
