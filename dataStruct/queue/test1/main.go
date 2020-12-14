/*
	数据结构
	队列 queue
	数组模拟队列（非环形队列）
*/
package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	queue   [5]int
	front   int
	rear    int
}

// 添加数据到队列
func (this *Queue) AddQueue(num int) (err error) {

	// 判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}

	this.rear++
	this.queue[this.rear] = num
	return

}

// 显示队列
func (this *Queue) ShowQueue() {

	fmt.Println("显示当前队列")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("queue[%d] = %d ", i, this.queue[i])
	}

}

// 出队列
func (this *Queue) GetQueue() (value int, err error) {

	// 先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	value = this.queue[this.front]
	return value, err
}

func main() {

	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var value int

	for {

		fmt.Println("\n1.输入 add 表示添加数据到队列")
		fmt.Println("2.输入 get 表示从队列中获取数据")
		fmt.Println("3.输入 show 表示显示队列")
		fmt.Println("4.输入 exit 表示退出程序")

		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("请输入要入队列的数：")
			fmt.Scanln(&value)
			err := queue.AddQueue(value)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("成功添加到队列")
			}
		case "get":
			value, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中取出的数为：", value)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}

	}

}
