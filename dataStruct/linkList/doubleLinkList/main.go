/*
	数据结构
	双向链表
*/
package main

import "fmt"

// 定义一个双向链表
type DoubleNode struct {
	no       int         // 排名
	name     string      // 名字
	nickname string      // 外号
	pre      *DoubleNode // 前一个节点
	next     *DoubleNode // 下一个节点
}

// 给双向链表加入一个节点
// 第 1 种插入方式：在双向链表的尾部加入节点
func InsertDoubleNode1(head *DoubleNode, newNode *DoubleNode) {

	// 先找到该链表的最后一个节点
	// 创建一个辅助节点
	// 辅助节点是用来帮助寻找单向链表的最后一个节点
	temp := head
	for temp.next != nil {
		temp = temp.next // 继续遍历下一个节点
	}

	// 找到最后一个节点后
	// 将 newNode 加入到最后一个节点
	temp.next = newNode
	// 将 newNode 的前一个节点指向 temp
	newNode.pre = temp

}

// 第 2 种插入方式：根据 no ,有序地插入节点
func InsertDoubleNode2(head *DoubleNode, newNode *DoubleNode) {

	temp := head
	var flag bool = true

	for {

		if temp.next == nil {
			break
		}

		if temp.next.no > newNode.no {
			break
		}

		if temp.next.no == newNode.no {
			flag = false
			break
		}

		temp = temp.next

	}

	if flag {

		newNode.next = temp.next
		newNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newNode
		}
		temp.next = newNode

	} else {
		fmt.Printf("位置 %d 已经存在节点", newNode.no)
	}

}

// 双向链表删除指定节点
func DelDoubleNode(head *DoubleNode, id int) {

	temp := head
	var flag bool = false

	for {

		if temp.next == nil {

			if temp.no == id {
				flag = true
			}

			break

		} else {

			if temp.no == id {
				flag = true
				break
			}

		}

		temp = temp.next

	}

	if flag {

		if temp.next != nil {
			temp.next.pre = temp.pre
		}
		temp.pre.next = temp.next

	} else {
		fmt.Println("\n双向链表中不存在该节点")
	}

}

// 遍历双向链表
func ListDoubleNode(head *DoubleNode) {

	temp := head

	// 先判断该链表是否为空
	if temp.next == nil {
		fmt.Println("empty linked list")
		return
	}

	fmt.Println("\n顺序的链表信息为：")
	// 循环遍历单向链表
	for {
		// 通过前一个节点的索引，打印后一个节点的内容
		fmt.Printf("[%d %s %s] ==> ", temp.next.no, temp.next.name,
			temp.next.nickname)

		temp = temp.next

		if temp.next == nil {
			break
		}

	}

}

// 逆序遍历双向链表
func ReverseDoubleNode(head *DoubleNode) {

	// 判断是否为空链表
	if head.next == nil {
		fmt.Println("empty list")
		return
	}

	// 寻找最后一个节点
	temp := head
	for temp.next != nil {
		temp = temp.next
	}

	// 逆序输出双向链表
	fmt.Println("\n逆序输出双向链表：")
	for temp.pre != nil {

		fmt.Printf("[%d %s %s] <== ", temp.no, temp.name,
			temp.nickname)

		temp = temp.pre

	}

}

func main() {

	// 先创建一个头节点
	head := &DoubleNode{}

	// 创建一个新的 DoubleNode
	node1 := &DoubleNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	node2 := &DoubleNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	node3 := &DoubleNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	node4 := &DoubleNode{
		no:       4,
		name:     "公孙胜",
		nickname: "入云龙",
	}

	// 无序地添加节点
	InsertDoubleNode1(head, node1)
	InsertDoubleNode1(head, node2)
	InsertDoubleNode1(head, node3)
	InsertDoubleNode1(head, node4)
	ListDoubleNode(head)

	// 逆序遍历双向链表
	ReverseDoubleNode(head)

	// 有序地添加节点
	head2 := &DoubleNode{}
	fmt.Println("有序地添加节点============")
	InsertDoubleNode2(head2, node3)
	InsertDoubleNode2(head2, node2)
	InsertDoubleNode2(head2, node4)
	InsertDoubleNode2(head2, node1)
	// 正序遍历
	ListDoubleNode(head2)
	// 反序遍历
	ReverseDoubleNode(head2)

	// 删除双向链表的某个节点
	DelDoubleNode(head2, 3)
	ListDoubleNode(head2)
	ReverseDoubleNode(head2)

}
