package main

import "fmt"

/*
	继承和接口的关系
*/

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbTree() {

	fmt.Println((*monkey).Name, "can climb tree")

}

// 声明接口
type BirdAble interface {
	Flying()
}

// 继承了 Monkey 的 LittleMonkey
type LittleMonkey struct {
	Monkey
}

// 让 LittleMonkey 实现 BirdAble 接口的 Flying 方法
func (littleMonkey *LittleMonkey) Flying() {
	fmt.Println((*littleMonkey).Name, "can flying")
}

func main() {

	littleMonkey := LittleMonkey{

		Monkey{
			Name: "Goku",
		},
	}

	(&littleMonkey).climbTree()
	(&littleMonkey).Flying()

}
