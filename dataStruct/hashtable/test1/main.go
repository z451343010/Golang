/*
	数据结构
	哈希表
*/
package main

import (
	"fmt"
	"os"
)

// 雇员
type Emp struct {
	Id   int    // 编号
	Name string // 姓名
	Next *Emp   // 指向下一个Emp的指针
}

// 雇员链表
// 不带表头，第一个节点直接存放数据
type EmpLink struct {
	Head *Emp // 该链表的第一个雇员
}

// 添加员工的方法
// 保证添加时，编号是从小到大
func (this *EmpLink) InsertEmp(emp *Emp) {

	cur := this.Head // 辅助指针

	// 如果当前的 EmpLink 是一个空链表
	if cur.Next == nil {
		cur.Next = emp
		return
	}

	flag := true

	for {

		if cur.Next == nil {
			break
		}

		if cur.Next.Id == emp.Id {
			flag = false
			break
		}

		if cur.Next.Id > emp.Id {
			break
		}

		cur = cur.Next

	}

	if flag {
		emp.Next = cur.Next
		cur.Next = emp
	} else {
		fmt.Printf("已经有了编号为 %d 的员工\n", emp.Id)
	}

}

// 显示当前链表的信息
func (this *EmpLink) ShowLink(no int) {

	if this.Head == nil || this.Head.Next == nil {
		fmt.Printf("当前链表 %d 为空", no)
		return
	}

	// 遍历当前的链表并显示数据
	cur := this.Head.Next // 辅助节点
	for cur != nil {
		fmt.Printf("链表[%d] 雇员id = %d 名字 = %s -> ", no, cur.Id, cur.Name)
		cur = cur.Next
	}

}

// 哈希表
// 本质是一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 散列函数
func (this *HashTable) HashFunc(id int) int {
	return id % len(this.LinkArr)
}

// Hash table 的 Insert 方法
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数
	// 确定将雇员添加到 Hash table 的那个链表中
	linkNo := this.HashFunc(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].InsertEmp(emp)

}

// 查找某个雇员
func (this *HashTable) FindEmp(id int) {

	linkNo := this.HashFunc(id)
	temp := this.LinkArr[linkNo].Head

	var flag bool = false
	for temp != nil {
		if temp.Id == id {
			fmt.Println("已经找到该雇员：")
			fmt.Printf("链表[%d] 雇员 id = %d 姓名 = %s\n",
				linkNo, temp.Id, temp.Name)
			flag = true
			break
		}
		temp = temp.Next
	}

	if !flag {
		fmt.Println("未找到该雇员")
	}

}

// 删除某个雇员
func (this *HashTable) DelNodeById(id int) {

	linkNo := this.HashFunc(id)
	temp := this.LinkArr[linkNo].Head
	flag := false

	// 如果对应的链表为空
	if temp.Next == nil {
		fmt.Println("该链表为空")
		return
	}

	for temp != nil {

		if temp.Next.Id == id {
			flag = true
			break
		}

		temp = temp.Next

	}

	if flag {
		fmt.Printf("节点 %d 已经被删除\n", id)
		temp.Next = temp.Next.Next
	} else {
		fmt.Printf("未找到节点 %d", id)
	}

}

// 显示 Hash table 的所有雇员
func (this *HashTable) ShowAllLink() {

	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
		fmt.Println()
	}

}

func main() {

	key := ""
	id := 0
	name := ""
	hashTable := &HashTable{}

	// 初始化 hashTable.LinkArr 里的 head 节点
	for i := 0; i < len(hashTable.LinkArr); i++ {
		head := &Emp{}
		hashTable.LinkArr[i].Head = head
	}

	for {
		fmt.Println("\n============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("delete 表示删除雇员")
		fmt.Println("exit 表示退出系统")

		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员的编号:")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员的名字:")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAllLink()
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			hashTable.FindEmp(id)
		case "delete":
			fmt.Println("请输入要删除的节点id：")
			fmt.Scanln(&id)
			hashTable.DelNodeById(id)
		case "exit":
			os.Exit(0)
		}

	}
}
