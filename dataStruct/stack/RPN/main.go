/*
	数据结构
	栈（Stack）
	逆波兰式（Reverse Polish notation）
	后缀表达式
	使用逆波兰式来简单模拟计算器的功能
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int     // 表示栈最大能存放的数据个数
	Top    int     // 栈顶
	arr    [50]int // 模拟栈的数组
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

// 判断一个字符是不是运算符[+, -, *, /]
func (this *Stack) IsOper(val int) bool {

	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}

}

// 运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// 返回运算符的优先级
func (this *Stack) Priority(oper int) int {

	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 47 {
		return 0
	} else {
		return -1
	}

}

func main() {

	// 数栈
	numStack := &Stack{
		MaxTop: 50,
		Top:    -1,
	}

	// 符号栈
	operStack := &Stack{
		MaxTop: 50,
		Top:    -1,
	}

	// 运算表达式
	// exp := "3+2*6-2"
	exp := "30+2*6-2"
	// 扫描表达式的索引
	index := 0

	// 定义运算的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {

		ch := exp[index : index+1]
		// 先把字符串转成 byte 切片
		// 再转成 int，字符对应的 ASCⅡ码值
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) {

			// 如果是空栈，直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)

			} else {

				// 判断新加入的符号和栈顶符号的运算优先级
				if operStack.Priority(operStack.arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					// 将计算后的结果重新压入数栈
					numStack.Push(result)
					// 把当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}

			}

		} else {

			keepNum += ch

			// 如果数字为多位数，则需要拼接
			// 每次都需要往 index 后面探一下，看看是不是运算符
			// 如果已经扫描到了最后一个字符
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向 index 后探一位
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}

			}

			// fmt.Println(keepNum)
			// val, _ := strconv.ParseInt(ch, 10, 64)
			// numStack.Push(int(val))
		}

		// 继续取下一个字符
		// 先判断 index 是否已经扫描到了最后一个字符
		if index == len(exp)-1 {
			break
		} else {
			index++
		}

	}

	// 如果表达式全部扫描完毕，依次从符号栈取出符号
	// 然后从数栈取出两个数，运算后的结果，入数栈，直到符号栈为空。
	for {

		if operStack.Top == -1 {
			break
		}

		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		// 将计算结果重新压入栈
		numStack.Push(result)

	}

	res, _ := numStack.Pop()
	fmt.Println("res = ", res)

}
