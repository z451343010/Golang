/*
	测试 Golang 指针变量的引用传递
*/
package main

import "fmt"

type SimpleList struct {
	num  int
	next *SimpleList
}

func main() {

	node3 := &SimpleList{
		num: 3,
	}
	node2 := &SimpleList{
		num:  2,
		next: node3,
	}
	node1 := &SimpleList{
		num:  1,
		next: node2,
	}
	head := &SimpleList{
		next: node1,
	}

	fmt.Println("head:", head)
	fmt.Println("node1:", node1)
	fmt.Println("node2:", node2)
	fmt.Println("node3:", node3)

	temp := head
	temp.next = temp.next.next
	fmt.Println("head.next:", head.next)

	fmt.Println("指针变量 head 的地址为：", &head)
	fmt.Println("指针变量 temp 的地址为：", &temp)
	fmt.Println("指针变量 head 的值为：", head)
	fmt.Println("指针变量 temp 的值为：", temp)

}
