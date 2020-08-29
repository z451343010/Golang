package main

import "fmt"

/*
	类型断言的最佳实践2
	判断输入的变量的类型
*/

type EmptyInterface interface {
}

type Student struct {
}

func TypeJudge(items ...interface{}) {

	for i, x := range items {

		switch x.(type) {

		case bool:
			fmt.Printf("param %d is a bool value is %v\n", i, x)
		case float64:
			fmt.Printf("param %d is a float64 value is %v\n", i, x)
		case int, int64:
			fmt.Printf("param %d is a int, int64 value is %v\n", i, x)
		case nil:
			fmt.Printf("param %d is a nil value is %v\n", i, x)
		case string:
			fmt.Printf("param %d is a string value is %v\n", i, x)
		case Student:
			fmt.Printf("param %d is a Student value is %v\n", i, x)
		case *Student:
			fmt.Printf("param %d is a *Student value is %v\n", i, x)
		default:
			fmt.Printf("param %d is unknown value is %v\n", i, x)

		}

	}

}

func main() {

	var a int64 = 15
	var b bool
	var c Student
	var d *Student
	var emptyInterface EmptyInterface
	emptyInterface = d

	_, flag := emptyInterface.(*Student)
	fmt.Println(flag)

	TypeJudge(a)
	TypeJudge(b)
	TypeJudge(c)
	TypeJudge(d)

}
