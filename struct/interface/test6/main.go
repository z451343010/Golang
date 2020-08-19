package main

import "fmt"

/*
	空接口interface{}没有任何方法
	所以所有类型都实现了空接口
	可以把任何一个变量赋值给一个空接口。
*/

type Empty interface {
}

type Student struct {
}

func main() {

	var stu Student
	var e Empty = stu

	fmt.Println(stu)
	fmt.Println(e)

}
