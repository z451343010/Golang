package main

import "fmt"

// Go的数组属于值类型。在默认情况下是值传递，因此会进行值拷贝。
// 数组间不会相互影响。（在很多其它编程语言中，数组是引用类型）
func test(testArr [3]int) {
	testArr[0] = 88
}

func main() {
	var arr = [3]int{11, 22, 33}
	fmt.Println(arr)
}
