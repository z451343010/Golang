package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (student *Student) say() string {

	// var returnStr string
	// returnStr = "Name:", (*student).Name, "Gender:",
	// 	(*student).Gender, "Age:",
	// 	(*student).Age,
	// 	"Id:", (*student).Id, "Score:", (*student).Score
	// returnStr := "Name:"
	// returnStr += (*student).Name
	// returnStr += " "
	// returnStr += "Gender:"
	// returnStr += (*student).Gender
	// returnStr += " "
	// returnStr += "Age:"
	// returnStr += strconv.FormatInt(int64((*student).Age), 10)
	// returnStr += " "
	// returnStr += "Id:"
	// returnStr += strconv.FormatInt(int64((*student).Id), 10)
	// returnStr += " "
	// returnStr += "Score:"
	// returnStr += strconv.FormatFloat((*student).Score,
	// 	'E', -1, 64)
	var returnStr = fmt.Sprintf("Name = %v Gender = %v Age = %v Id = %v Score = %v",
		student.Name, student.Gender, student.Age, student.Id, student.Score)

	return returnStr

}

func main() {

	// var stu = Student{
	// 	Name:   "zjune",
	// 	Gender: "male",
	// 	Age:    24,
	// 	Id:     2014211,
	// 	Score:  100,
	// }

	var student Student
	student.Name = "zjune"
	student.Gender = "male"
	student.Age = 24
	student.Id = 2014211
	student.Score = 100

	fmt.Println((&student).say())

}
