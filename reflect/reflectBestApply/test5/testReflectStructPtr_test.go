/*
	使用反射创建并操作结构体
*/
package test5

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {

	// 一次性定义多个变量
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model) // 获取类型 *user
	t.Log("reflect.TypeOf", st.Kind().String())
	st = st.Elem()
	t.Log("reflect.TypeOf.Elem", st.Kind().String())
	// New 返回一个Value类型值
	// 该值持有一个指向类型为typ的新申请的零值的指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String())
	// model就是创建的user结构体变量（实例）
	// 先转成空接口，再进行类型断言(*user)
	model = elem.Interface().(*user)
	elem = elem.Elem() // 取得elem指向的值
	elem.FieldByName("UserId").SetString("12345678")
	elem.FieldByName("Name").SetString("nickname")
	t.Log("model model.Name", model, model.Name)
}
