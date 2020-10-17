/*
	请编写一个案例，演示对（基本数据类型、interface{}、
	reflect.Value）进行反射的基本操作。
*/

package main

import (
	"fmt"
	"reflect"
)

// 使用 interface{} 类型的参数
// 可以接收任何数据类型
func reflectTest(b interface{}) {

	// 通过反射获取变量的类型、类别、值
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	// 获取 reflect.Value
	// reValue 是一个 reflect.Value类型，而不是int类型
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue = ", rValue)
	fmt.Printf("rValue type = %T\n", rValue)
	// 将 reflect.Value类型转化成 int 类型
	num2 := 100 + rValue.Int()
	fmt.Printf("type of rValue.Int() = %T\n", rValue.Int())
	fmt.Println(num2)

	// 将 rValue 转成 interface{}
	iV := rValue.Interface()
	fmt.Println(iV)

	// 将  interface{} 通过断言转成需要的类型
	num3 := iV.(int)
	fmt.Println("num3 = ", num3)

}

func main() {

	var num int = 100
	reflectTest(num)

}
