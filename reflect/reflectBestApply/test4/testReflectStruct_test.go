/*
	使用反射操作任意结构体类型
*/
package main

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {

	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickname")
	t.Log("model", model)

}
