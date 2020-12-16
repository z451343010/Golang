/*
	使用带head头节点的单向链表实现-水浒传英雄排行榜管理：
	1）完成对英雄人物的增删改查操作。
	2）第一种方法：在添加英雄时，直接添加到链表的尾部
	3）第二种方法：在添加英雄时，根据排名将英雄插入到指定位置
	（如果没有这个排名，则添加失败，并给出提示）
*/
package main

import "fmt"

// 定义一个 HeroNode
type HeroNode struct {
	no       int       // 排名
	name     string    // 名字
	nickname string    // 外号
	next     *HeroNode // 下一个节点
}

// 给链表加入一个节点
// 第 1 种插入方式：在单链表的最后插入节点
func InsertHeroNode1(head *HeroNode, newHeroNode *HeroNode) {

	// 先找到该链表的最后一个节点
	// 创建一个辅助节点
	// 辅助节点是用来帮助寻找单向链表的最后一个节点
	temp := head
	for temp.next != nil {
		temp = temp.next // 继续遍历下一个节点
	}

	// 将 newHeroNode 加入到最后一个节点
	temp.next = newHeroNode

}

// 给链表加入一个节点
// 第 2 种插入方式：根据 no ，有序地插入链表
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {

	// 先找到链表中适当的节点
	// 创建一个辅助节点
	// 辅助节点是用来帮助寻找单向链表适当的节点
	temp := head
	var flag bool = true

	// 循环寻找适合插入的节点
	for {

		if temp.next == nil {
			break
		}
		if temp.next.no > newHeroNode.no {
			break
		}

		if temp.next.no == newHeroNode.no {
			flag = false
			break
		}

		// 如果未找到合适的节点，则继续尝试下一个节点
		temp = temp.next

	}

	if flag {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	} else {
		fmt.Printf(" %d 位置已经有节点\n", temp.next.no)
	}

}

// 删除单向链表的节点
func DelHeroNode(head *HeroNode, id int) {

	temp := head
	var flag bool = false // 是否找得到该id的节点

	for {

		if temp.next == nil {
			break
		}

		if temp.next.no == id {
			flag = true
			break
		}

		temp = temp.next

	}

	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("can not find the node")
	}

}

// 遍历单向链表
func ListHeroNode(head *HeroNode) {

	temp := head

	// 先判断该链表是否为空
	if temp.next == nil {
		fmt.Println("empty linked list")
		return
	}

	fmt.Println("链表信息为：")
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

func main() {

	// 先创建一个头节点
	head := &HeroNode{}

	// 为第二种插入方式创建头节点 head2
	head2 := &HeroNode{}

	// 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "公孙胜",
		nickname: "入云龙",
	}

	InsertHeroNode1(head, hero1)
	InsertHeroNode1(head, hero2)
	InsertHeroNode1(head, hero3)
	InsertHeroNode1(head, hero4)

	ListHeroNode(head)

	fmt.Println("\n有序地插入节点============")
	InsertHeroNode2(head2, hero3)
	InsertHeroNode2(head2, hero2)
	InsertHeroNode2(head2, hero1)
	InsertHeroNode2(head2, hero4)

	// 设置 0 号变量，查看是否会按顺序
	hero0 := &HeroNode{
		no:       0,
		name:     "鲁路修",
		nickname: "ZERO",
	}
	InsertHeroNode2(head2, hero0)

	// 设置 no 重复的变量测试
	hero33 := &HeroNode{
		no:       3,
		name:     "有用",
		nickname: "智少星",
	}
	InsertHeroNode2(head2, hero33)

	// 设置 99 号变量，测试
	hero99 := &HeroNode{
		no:       99,
		name:     "刹那 F 清英",
		nickname: "能天使",
	}
	InsertHeroNode2(head2, hero99)

	ListHeroNode(head2)

	DelHeroNode(head2, 0)
	fmt.Println("\n删除了节点后的链表为============")
	ListHeroNode(head2)

}
