/*
	反射的最佳应用实践
*/
package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {

	fmt.Println("---Print start---")
	fmt.Println(s)
	fmt.Println("---Print end---")

}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int,
	score float32, sex string) {

	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex

}

func TestStruct(a interface{}) {

	// 获取 reflect.Type 类型
	typ := reflect.TypeOf(a)
	// 获取 reflect.Value
	val := reflect.ValueOf(a)
	// 获取 a 对应的类别
	kd := val.Kind()
	// 判断 a 的类别是不是结构体 struct
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	// 获取该结构体一共有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		// 获取到struct标签，注意需要通过reflect.Type来获取Tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//  该字段有 tag 标签，才会显示
		if tagVal != "" {
			fmt.Printf("Field %d:tag = %v\n", i, tagVal)
		}
	}

	// 获取结构体有多少个方法 method
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// 获取到第二个方法
	// Call 代表调用该方法
	// 因为方法的排序默认是按照方法名首字母的 ASCⅡ码
	// 因此顺序为 GetSum(G:0) Print(P:1) Set(S:2)
	val.Method(1).Call(nil) // Print()

	// 调用结构体的第1个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	// 传入的参数是 []reflect.Value
	res := val.Method(0).Call(params) // GetSum()
	// 返回结果，返回的结果是 []reflect.Value
	fmt.Println("res = ", res[0].Int())

}

func main() {

	// 创建了一个 Monster 实例
	var a Monster = Monster{
		Name:  "Diablo",
		Age:   4000,
		Score: 100.0,
		Sex:   "male",
	}

	// 将 Monster 实例传递给 TestStruct 函数
	TestStruct(a)

}
