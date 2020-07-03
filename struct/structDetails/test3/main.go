package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 重新定义
type Stu Student

// 结构体进行type重新定义（相当于取别名）
// Golang认为是新的数据类型，但是相互间可以强制转换
func main() {

	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1)

	fmt.Println(stu1)
	fmt.Println(stu2)

}
