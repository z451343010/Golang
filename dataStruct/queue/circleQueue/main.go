/*
	数据结构
	环形队列
*/
package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // 指向队列队首
	tail    int // 指向队列队尾

}

// 入队列
func (this *CircleQueue) Push(value int) (err error) {

	if this.IsFull() {
		return errors.New("queue full")
	}

	// this.tail 在队列尾部，但是不包含队列最后的元素
	this.array[this.tail] = value // 把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队列
func (this *CircleQueue) Pop() (value int, err error) {

	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	// 取出 head
	// head 指向队首，并且包含队首元素
	value = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return value, err

}

// 显示队列
func (this *CircleQueue) ShowQueue() {

	fmt.Println("环形队列情况如下：")
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	// 设计一个辅助变量，指向 head
	tempHead := this.head

	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d ", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}

	// 换行
	fmt.Println()

}

// 判断环形队列是否已经满了
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {

	// 初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var value int

	for {

		fmt.Println("\n1.输入 push 表示添加数据到环形队列")
		fmt.Println("2.输入 pop 表示从环形队列中获取数据")
		fmt.Println("3.输入 show 表示显示环形队列")
		fmt.Println("4.输入 num 表示显示队列元素的个数")
		fmt.Println("5.输入 exit 表示退出程序")
		fmt.Println("queue.head -> ", queue.head)
		fmt.Println("queue.tail -> ", queue.tail)

		fmt.Scanln(&key)

		switch key {
		case "push":
			fmt.Println("请输入要入队列的数：")
			fmt.Scanln(&value)
			err := queue.Push(value)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("成功添加到队列")
			}
		case "pop":
			value, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中取出的数为：", value)
			}
		case "show":
			queue.ShowQueue()
		case "num":
			fmt.Println("环形队列中的元素个数是：", queue.Size())
		case "exit":
			os.Exit(0)
		}

	}

}
