/*
	对结构体的反射
*/
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	// 获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	// 获取到 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue = ", rValue)
	fmt.Printf("type rValue = %T\n", rValue)

	// 获取 rValue 对应的 Kind
	fmt.Println("Kind of rValue = ", rValue.Kind())
	fmt.Println("Kind of rValue = ", rType.Kind())

	// 将 rValue 转成 interface{}
	iV := rValue.Interface()
	fmt.Printf("type of iV = %T\n", iV)

	// 将 interface{} 类型的变量
	// 通过类型断言转化成需要的类型
	stu := iV.(Student)
	fmt.Println("stu = ", stu)
	fmt.Printf("type stu = %T\n", stu)

}

type Student struct {
	Name string
	Age  int
}

func main() {

	stu := Student{
		Name: "zhang",
		Age:  20,
	}

	reflectTest(stu)

}
