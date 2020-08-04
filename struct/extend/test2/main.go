package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

// 将 Pupil 和 Graduate 共有的方法也绑定到 Student 类型
func (student *Student) ShowInfo() {
	fmt.Printf("学生姓名 = %v 年龄 = %v 成绩 = %v\n",
		(*student).Name, (*student).Age, (*student).Score)
}

func (student *Student) SetScore(score int) {
	(*student).Score = score
}

// 小学生
type Pupil struct {
	// 嵌入 Student 的匿名结构体
	Student
}

// 当我们对结构体嵌入了匿名结构体，使用方式会发生变化
func main() {

	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 18
	pupil.Student.SetScore(99)
	pupil.ShowInfo()

}
