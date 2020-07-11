package main

import "fmt"

// 如果一个变量实现了String()这个方法
// 那么fmt.Println()会默认调用这个变量的String()进行输出

type Test struct {
}

func (test *Test) String() string {

	return "test String"

}

func main() {

	var test Test
	// 如果实现了 *Test 类型的 String 方法，输出时就会自动调用
	fmt.Println(&test)

}
