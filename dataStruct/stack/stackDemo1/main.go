/*
	数据结构
	栈（Stack）
	快速入门案例
*/
package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int    // 表示栈最大能存放的数据个数
	Top    int    // 栈顶
	arr    [5]int // 模拟栈的数组
}

// 入栈
func (this *Stack) Push(val int) (err error) {

	// 先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("push stack error,stack full")
		return errors.New("push stack error,stack full")
	}

	this.Top++
	this.arr[this.Top] = val
	return

}

// 出栈
func (this *Stack) Pop() (val int, err error) {

	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("pop error, stack is empty")
		return 0, errors.New("pop error, stack is empty")
	}

	val = this.arr[this.Top]
	this.Top--
	return val, nil

}

// 遍历栈
func (this *Stack) ShowStack() {

	fmt.Println("栈的情况如下：")

	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("empty stack")
		return
	}

	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d\n", i, this.arr[i])
	}

}

func main() {

	// 初始化栈
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}

	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	// 显示
	stack.ShowStack()

	// 出栈
	val, err := stack.Pop()
	if err == nil {
		fmt.Println("出栈的值为：", val)
	}

	// 显示
	stack.ShowStack()

}
