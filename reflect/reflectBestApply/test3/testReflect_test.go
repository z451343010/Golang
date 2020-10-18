/*
	定义了两个函数 test1 test2
	定义一个适配器函数用作统一处理接口
*/
package test3

import (
	"reflect"
	"testing"
)

func TestReflectFunc(t *testing.T) {

	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	// 适配器
	// call 传入函数
	// args 传入数量可变的参数
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}

		function = reflect.ValueOf(call)
		function.Call(inValue)

	}
	var v1 int = 10
	var v2 int = 20
	var str string = "zhang"
	bridge(call1, v1, v2)
	bridge(call2, v1, v2, str)

}
