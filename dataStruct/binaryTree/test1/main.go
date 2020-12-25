/*
	数据结构
	二叉树
*/
package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍历
// 先输出根节点，再输出左子树，右子树
func PreOrder(node *Hero) {

	if node != nil {

		fmt.Printf("no = %d name = %s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)

	}

}

// 中序遍历
// 先输出左子树、根节点、右子树
func InfixOrder(node *Hero) {

	if node != nil {

		InfixOrder(node.Left)
		fmt.Printf("no = %d name = %s\n", node.No, node.Name)
		InfixOrder(node.Right)

	}

}

// 后序遍历
// 先输出左子树、右子树、根节点
func PostOrder(node *Hero) {

	if node != nil {

		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no = %d name = %s\n", node.No, node.Name)

	}

}

func main() {

	// 构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}

	left1 := &Hero{
		No:   2,
		Name: "卢俊义",
	}

	right1 := &Hero{
		No:   3,
		Name: "吴用",
	}

	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No:   4,
		Name: "公孙胜",
	}

	right1.Right = right2

	fmt.Println("========== 前序遍历 ==========")
	PreOrder(root)

	fmt.Println("========== 中序遍历 ==========")
	InfixOrder(root)

	fmt.Println("========== 后序遍历 ==========")
	PostOrder(root)

}
