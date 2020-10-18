/*
	给定一个变量 var v float64 = 1.2
	用反射来得到它的 reflect.Value,然后获取对应的Type
	Kind 和值，并将reflect.Value转换成interface{}
	再将interface{}转换成float64
*/
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	rValue := reflect.ValueOf(b)
	fmt.Println("rValue = ", rValue)

	rType := reflect.TypeOf(rValue)
	fmt.Printf("rType = %v Type = %T\n", rType, rType)

	rKind := rType.Kind()
	fmt.Println("rKind = ", rKind)

	iValue := rValue.Interface()
	num := iValue.(float64)
	fmt.Printf("num = %v, type = %T\n", num, num)

}

func main() {

	var num float64 = 1.2
	reflectTest(num)

}
