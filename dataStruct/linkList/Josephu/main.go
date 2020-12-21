/*
	数据结构
	约瑟夫问题（Josephus problem）：
	1）设编号为 1，2，3 ... n 的 n 个人围坐一圈
	2）约定编号为 k （1 <= k <= n）的人从 1 开始报数
	数到 m 的那个人出列
	3）它的下一位又从 1 开始报数，数到 m 的那个人又出列
	依此类推，直到所有人出列为止
	4）由此产生一个出列编号的序列
*/
package main

import "fmt"

// 小孩的结构体
type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针
}

// 小孩构成单向环形链表
// num：小孩的个数
// *Boy：返回单向环形链表的头指针
func AddBoy(num int) *Boy {

	// 第一个小孩，空节点
	first := &Boy{}

	// 因为 first 指针固定指向第一个小孩
	// 因此 first 指针不能移动，需要一个辅助指针
	// 辅助指针 curBoy ,指向当前的小孩
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num 的值有误")
		return first
	}

	// 循环地构成单向环形链表
	for i := 1; i <= num; i++ {

		boy := &Boy{
			No: i,
		}

		// 第一个小孩
		if i == 1 {
			first = boy
			curBoy = boy
			// 形成一个自循环
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			// 最后一个节点指向头节点
			// 形成一个单向环形链表
			curBoy.Next = first
		}

	}

	return first

}

// 显示单向的环形链表
func ShowBoy(first *Boy) {

	// 单独处理空链表
	if first.Next == nil {
		fmt.Println("empty link list")
		return
	}

	// 创建一个辅助指针，帮助遍历
	curBoy := first
	for {

		fmt.Printf("小孩编号 = %d -> ", curBoy.No)
		// 退出循环的条件
		// 单向环形链表的最后一个节点指向头指针
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}

}

// 统计小孩的数量
func CountBoy(first *Boy) int {

	// 单独处理空链表
	if first.Next == nil {
		return 0
	}

	// 创建一个辅助指针，帮助遍历
	curBoy := first
	// 计数
	count := 0
	for {
		count++
		// 退出循环的条件
		// 单向环形链表的最后一个节点指向头指针
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}

	return count

}

// 解决约瑟夫问题的算法
// k : 约定编号为 k 的人从 1 开始报数
// m : 从1开始报数，数到 m 的人出列
func Josephu(first *Boy, k int, m int) {

	fmt.Println("\n小孩出圈的顺序如下：")

	// 单独处理空链表
	if first.Next == nil {
		fmt.Println("empty link list")
		return
	}

	if k > CountBoy(first) {
		fmt.Printf("输入的参数 %d 有误\n", k)
		return
	}

	// 定义一个辅助指针，帮助删除小孩
	tail := first
	// 让 tail 指向环形链表的最后一个节点（小孩）
	// 后面 tail 和 first 都往后挪动的时候
	// 才能保证 tail 一直在 fist 后面
	for {
		// 说明已经指向了最后一个节点
		// 退出循环
		if tail.Next == first {
			break
		}

		tail = tail.Next

	}

	// 让 first 移动到 k （编号为 k 的人开始报数）
	// 移动到编号 k ,只需要移动 k - 1 次
	for i := 0; i < k-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 开始数 m 下（数到 m 的人出列）
	// 然后就删除 first 指向的小孩
	for {

		for i := 0; i < m-1; i++ {
			first = first.Next
			tail = tail.Next
		}

		fmt.Printf("编号为 %d 的小孩出圈\n", first.No)
		// 删除 first 指向的小孩
		first = first.Next
		tail.Next = first

		// 如果圈中只有一个小孩
		if first == tail {
			fmt.Printf("编号为 %d 的小孩出圈\n", first.No)
			break
		}
	}

}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
	Josephu(first, 2, 3)
}
