/*
	使用反射的方式来获取结构体的tag标签，遍历字段的值，修改字段值
	调用结构体方法（要求：通过传递地址的方式完成）
*/
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"monster_name"`
	Age   int
	Score float32
	Sex   string
}

func (s Monster) Print() {

	fmt.Println("---Print start---")
	fmt.Println(s)
	fmt.Println("---Print end---")

}

func TestStruct(a interface{}) {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("白象精")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n", num)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag = %s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	val.Elem().Method(0).Call(nil)

}

func main() {

	var a Monster = Monster{
		Name:  "YellowLion",
		Age:   408,
		Score: 92.8,
	}

	// Marshal 就是通过反射获取到 struct 的 tag 值
	result, _ := json.Marshal(a)
	fmt.Println("json result :", string(result))

	TestStruct(&a)
	fmt.Println(a)

}
