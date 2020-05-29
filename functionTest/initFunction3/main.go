package main

import (
	"fmt"
	"functionTest/initFunction3/students"
)

var test int = testFunc()

func testFunc() int {

	fmt.Println("main testFunc")
	return 10

}

func main() {

	var mainAge int8 = students.StuAge
	fmt.Println("mainAge = ", mainAge)
	// fmt.Println("StuAge =", students.StuAge)
	// fmt.Println("StuName =", students.StuName)

}
