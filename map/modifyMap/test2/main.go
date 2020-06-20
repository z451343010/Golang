package main

import "fmt"

// 定义结构体Student
type Student struct {
	Name    string
	Age     int
	Address string
}

// map的value也经常使用struct类型
// 更适合管理负责的数据（比前面的value是一个map更好）
// 比如value为Student结构体
// 1.map的key为学生的编号，是唯一的
// 2.map的value为Student结构体，包含学生的名字、年龄、地址
func main() {

	students := make(map[string]Student)
	student1 := Student{"zjune", 18, "杭州"}
	student2 := Student{"xanxus", 18, "杭州"}
	student3 := Student{"kkk", 18, "杭州"}

	students["no1"] = student1
	students["no2"] = student2
	students["no3"] = student3

	// fmt.Println(students)

	// 遍历map中的student
	for key, value := range students {
		fmt.Println("学生的编号是：", key)
		fmt.Println("学生的名字是：", value.Name)
		fmt.Println("学生的年龄是：", value.Age)
		fmt.Println("学生的地址是：", value.Address)
		fmt.Println()
	}

}
