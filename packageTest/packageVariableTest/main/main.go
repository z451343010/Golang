package main

import (
	"fmt"
	// 引入时写入路径
	test "packageTest/packageVariableTest/utils"
)

func main() {
	// 调用包里面的方法时，使用包名 test
	fmt.Println(test.Calculate())
}
