package main

import "fmt"

// 数组定义的四种方式
func main() {

	var numsArray1 [3]int = [3]int{1, 2, 3}
	var numsArray2 = [3]int{1, 2, 3}
	// 不声明数组的大小
	var numsArray3 = [...]int{6, 7, 8}
	// 可以指定元素值对应的下标
	var names = [3]string{1: "tome", 0: "jack", 2: "marry"}

	fmt.Println(numsArray1)
	fmt.Println(numsArray2)
	fmt.Println(numsArray3)
	fmt.Println(names)

}
