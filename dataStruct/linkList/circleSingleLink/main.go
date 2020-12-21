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
// func InsertCircleNodeByOrder(head *CatNode, newNode *CatNode) {

// 如果单向环形链表没有任何节点
// 	if head.next == nil {

// 		head.no = newNode.no
// 		head.name = newNode.name
// 		head.next = head

// 		return

// 	}

// 	temp := head

// 	for {

// 		if temp.next == head {
// 			if temp.no > newNode.no {

// 				newNode.next = temp.next
// 				newNode = head
// 				temp.next = newNode

// 			}
// 			break
// 		}

// 	}

// }

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
			fmt.Printf("删除的猫 = %d\n", id)
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

// 显示环形链表的节点
func ListCircleNode(head *CatNode) {

	temp := head

	for {
		fmt.Printf("[%d] %s => ", temp.no, temp.name)
		if temp.next == head {
			break
		}

		temp = temp.next

	}

	// fmt.Println("\ntemp => ", temp)
	// fmt.Println("temp.next => ", temp.next)

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

	// 测试有序排列数组
	// head2 := &CatNode{}

	// node4 := &CatNode{
	// 	no:   4,
	// 	name: "Cat4",
	// }

	// node5 := &CatNode{
	// 	no:   5,
	// 	name: "Cat5",
	// }

	// node33 := &CatNode{
	// 	no:   3,
	// 	name: "Cat33",
	// }

	fmt.Println("\n有序地显示节点============")
	// InsertCircleNodeByOrder(head2, node5)
	// InsertCircleNodeByOrder(head2, node4)
	// InsertCircleNodeByOrder(head2, node33)
	// InsertCircleNodeByOrder(head2, node6)
	// InsertCircleNodeByOrder(head2, node2)
	// ListCircleNode(head2)

}
