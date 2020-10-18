/*
	反射最佳实践：课堂练习
	要求
	1）编写一个Cal结构体，有两个字段Num1,Num2
	2）方法 GetSub(name string)
	3）使用反射遍历 Cal 结构体所有的字段信息
	4）使用反射机制完成对GetSub的调用，输出形式为
	“tom 完成了减法运行，8 - 3 = 5” （根据传入的参数输出）
*/
package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal) GetSub(name string) {
	result := cal.Num1 - cal.Num2
	fmt.Printf("%s 完成了减法运行，%d - %d = %d",
		name, cal.Num1, cal.Num2, result)
}

func TestReflect(b interface{}) {

	// typ := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	kd := rValue.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := rValue.NumField()
	fmt.Println("结构体的字段数量为：", num)

	for i := 0; i < num; i++ {
		fmt.Printf("字段 %d 的值为：%v\n", i, rValue.Field(i))
	}

	// 设置调用参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))

	// 使用反射机制完成对 GetSub 的调用
	rValue.Method(0).Call(params)

}

func main() {

	var cal = Cal{
		Num1: 8,
		Num2: 3,
	}

	TestReflect(cal)

}
