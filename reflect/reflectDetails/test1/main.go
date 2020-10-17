/*
	通过反射修改 num int 的值
	修改 student 的值
*/
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	// 获取到 reflect.Value
	rValue := reflect.ValueOf(b)
	// 查看 rValue 的 Kind
	fmt.Println("rValue Kind = ", rValue.Kind())
	// Elem() 返回 v 持有的接口保管的值的
	// Value封装，或者v持有的指针指向的值的Value封装
	// 因为传入的类型是指针类型，因此需要使用Elem()获取指针指向的值
	rValue.Elem().SetInt(20)

}

func main() {

	var num int = 10
	reflectTest(&num)
	fmt.Println("num = ", num)

}
