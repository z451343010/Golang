package main

import (
	"testing"
)

// 编写测试用例，看刚才的函数是否正确
func TestAddUpper(t *testing.T) {

	res := addUpper(10)
	if res != 55 {
		t.Fatalf("函数执行错误,函数期望值 = %v,实际值 = %v", 55, res)
		// t.Fatalf("函数期望值 = %v，实际值 = %v", 55, res)
	}

	// 如果执行正确，输出日志
	t.Logf("函数执行正确")

}
