/*
	数据结构
	单向环形链表
*/
package main

import "fmt"

type CatNode struct {
	no   int      // 猫编号
	name string   // 猫的名字
	next *CatNode // 下一个猫
}

// 无序地往环形链表添加节点
func InsertCircleNode(head *CatNode, newNode *CatNode) {

	// 判断进来的新节点，是不是第一个节点
	// 由于单向环形链表没有 head 节点
	// 因此默认第一个进来的节点为头节点
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}

	// 进来的新节点，不是第一个节点
	temp := head

	// 寻找最后一个节点
	for temp.next != head {
		temp = temp.next
	}

	newNode.next = head
	temp.next = newNode

}

// 按照 no, 有序地往环形链表中加入节点
func InsertCircleNodeByOrder(head *CatNode, newNode *CatNode) *CatNode {

	// 辅助节点
	temp := head
	// 辅助指针，永远指向单向环形队列的尾部
	tail := head

	// 新节点进来时，单向环形链表为空
	// 则让第一个节点为头节点，构成自循环
	if head.next == nil {
		head.name = newNode.name
		head.no = newNode.no
		head.next = head
		return head
	}

	// 新节点进来时，单向环形链表中只有一个自循环的节点
	if temp.next == head {

		head.next = newNode
		newNode.next = head

		// 比较 newNode 和头节点的id大小
		// 如果新节点的id较小，则称为头节点
		if newNode.no < temp.no {
			head = newNode
		}

	} else { // 单向环形链表有两个及以上的节点

		// 让 tail 指向尾部节点
		for tail.next != head {
			tail = tail.next
		}

		// 如果新节点的 no 小于头节点
		// 让新节点成为头节点
		if head.no > newNode.no {
			newNode.next = head
			tail.next = newNode
			head = newNode
			return head
		}

		// 如果新节点大于头节点
		// 循环地在单向环形链表中找到适合插入新节点的位置
		flag := false
		for {

			if temp.no == newNode.no {
				break
			}

			// 如果循环到了单向环形链表的尾部
			// 新节点的 no 大于单向环形链表的任何一个节点
			// 因此把新节点插入到单向环形链表的尾部
			if temp.next == head {
				flag = true
				break
			}

			if temp.next.no > newNode.no {
				flag = true
				break
			}

			temp = temp.next

		}

		if flag {

			newNode.next = temp.next
			temp.next = newNode

		} else {
			fmt.Println("编号为 %d 的节点已经存在", newNode.no)
		}

	}

	return head

}

// 删除环形单向链表的节点
func DelCircleNode(head *CatNode, id int) *CatNode {

	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("the link list is empty")
		return head
	}

	// 假设环形链表就只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}

	// 将 helper 定位到环形链表的最后
	for {

		if helper.next == head {
			break
		}

		helper = helper.next

	}

	// 链表有两个以上的节点
	flag := true
	for {

		// 说明到了最后一个节点，还没有找到删除的节点
		if temp.next == head {
			break
		}

		if temp.no == id {
			// 如果删除的是头节点
			if temp == head {
				head = head.next

			}
			helper.next = temp.next
			fmt.Printf("\n删除的猫 = %d\n", id)
			flag = false
			break
		}

		// 移动，用来比较 id 和节点的 no
		temp = temp.next
		// 移动，一旦找到要删除的节点
		helper = helper.next

	}

	// 这里还需要再比较一次
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("删除的猫 = %d\n", id)
		} else {
			fmt.Printf("没有找到需要删除的节点：%d\n", id)
		}
	}

	return head

}

// 显示单向环形链表的节点
func ListCircleNode(head *CatNode) {

	temp := head

	for {
		fmt.Printf("[%d] %s => ", temp.no, temp.name)
		if temp.next == head {
			break
		}

		temp = temp.next

	}

}

func main() {

	head := &CatNode{}

	node1 := &CatNode{
		no:   1,
		name: "Carfield",
	}

	node2 := &CatNode{
		no:   2,
		name: "American Shorthair",
	}

	node3 := &CatNode{
		no:   3,
		name: "Bule Cat",
	}

	InsertCircleNode(head, node1)
	InsertCircleNode(head, node2)
	InsertCircleNode(head, node3)
	head = DelCircleNode(head, 1)
	ListCircleNode(head)

	// 测试有序插入新节点
	head2 := &CatNode{}

	node4 := &CatNode{
		no:   4,
		name: "Cat4",
	}

	node5 := &CatNode{
		no:   5,
		name: "Cat5",
	}

	node33 := &CatNode{
		no:   3,
		name: "Cat33",
	}

	node66 := &CatNode{
		no:   6,
		name: "Cat66",
	}

	node88 := &CatNode{
		no:   88,
		name: "Cat88",
	}

	node7 := &CatNode{
		no:   7,
		name: "Cat7",
	}

	node77 := &CatNode{
		no:   7,
		name: "Cat77",
	}

	node22 := &CatNode{
		no:   2,
		name: "Cat22",
	}

	fmt.Println("\n有序地显示节点============")
	head2 = InsertCircleNodeByOrder(head2, node5)
	head2 = InsertCircleNodeByOrder(head2, node4)
	head2 = InsertCircleNodeByOrder(head2, node33)
	head2 = InsertCircleNodeByOrder(head2, node66)
	head2 = InsertCircleNodeByOrder(head2, node88)
	head2 = InsertCircleNodeByOrder(head2, node7)
	head2 = InsertCircleNodeByOrder(head2, node77)
	head2 = InsertCircleNodeByOrder(head2, node22)
	ListCircleNode(head2)

	// 测试删除有序环形链表
	head2 = DelCircleNode(head2, 6)
	ListCircleNode(head2)

	// 测试删除有序环形链表
	head2 = DelCircleNode(head2, 2)
	ListCircleNode(head2)

}
