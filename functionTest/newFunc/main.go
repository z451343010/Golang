package main

import "fmt"

func main() {

	num := new(int) // 分配的为指针类型 *int
	fmt.Printf("num的类型%T，num的值%v，num的地址%v\n",
		num, num, &num)

	//
	*num = 100
	fmt.Printf("num的类型%T，num的值%v，num的地址%v，指向的值为%v",
		num, num, &num, *num)

}
