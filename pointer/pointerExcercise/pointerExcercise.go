package main

import "fmt"

// 课堂练习
// 1.写一个程序获取一个int变量num的地址，并显示到终端
// 2.将num的地址赋值给指针ptr，并通过ptr去修改num的值
func main() {

	// 1
	var num int
	fmt.Println("the value of num is:", &num)
	fmt.Println("the address of num is:", num)

	// 2
	var ptr *int = &num
	*ptr = 3 // *ptr访问到了这个内存空间，重新给这个内存空间赋值
	fmt.Println("the changed value of num is:", num)
}
