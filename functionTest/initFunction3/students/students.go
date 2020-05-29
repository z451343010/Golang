package students

import "fmt"

var StuAge int8
var StuName string

var test int = testFunc()

func testFunc() int {

	fmt.Println("students testFunc")
	return 9

}

func init() {

	StuAge = 24
	StuName = "zhang"

	fmt.Println("students init function")

}
