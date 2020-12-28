/*
	数据结构
	队列 链表
	创建一个链表模拟队列
	实现：数据入队列、数据出队列、显示队列
*/
package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Num   int
	Next  *Node
	front *Node
	rear  *Node
}

// 数据入队列
func (this *Node) Insert(head *Node, num int) {

	newNode := &Node{
		Num: num,
	}

	// 如果一开始为空链表
	if head.Next == nil {
		head.Next = newNode
		this.rear = head
		this.rear = this.rear.Next
		return
	}

	// 如果不是空链表
	this.rear.Next = newNode
	this.rear = this.rear.Next

}

// 数据出队列
func (this *Node) Pop(head *Node) (num int, err error) {

	if head.Next == nil {
		return -1, errors.New("queue is empty")
	}

	this.front = head.Next

	// 如果只有一个节点
	if this.front == this.rear {
		head.Next = nil
		return this.front.Num, nil
	}

	// 不止一个节点
	head.Next = this.front.Next
	return this.front.Num, nil

}

// 显示数据队列
func (this *Node) Show(head *Node) {

	fmt.Println("\n队列中的数据如下:")

	// 每次都让 temp 指针指向 head 节点的下一个节点
	temp := head.Next

	// 如果队列中没有数据
	if temp == nil {
		fmt.Println("queue is empty")
		return
	}

	// 循环显示队列中的每个节点
	for {

		fmt.Printf("%d -> ", temp.Num)

		// 如果 temp 指针移动到了和尾部指针 rear 重合
		// 说明整个队列已经扫描完毕，退出循环
		if temp == this.rear {
			break
		}

		// 继续扫描队列中的下一个节点
		temp = temp.Next

	}

}

func main() {

	head := &Node{}
	queue := &Node{}

	// 入队列
	queue.Insert(head, 1)
	queue.Insert(head, 2)
	queue.Insert(head, 3)
	queue.Insert(head, 4)
	queue.Insert(head, 5)

	// 显示队列
	queue.Show(head)

	// 出队列
	var num int
	num, _ = queue.Pop(head)
	fmt.Println("\n出队列的数据为：", num)
	num, _ = queue.Pop(head)
	fmt.Println("\n出队列的数据为：", num)

	// 显示队列
	queue.Show(head)

}
