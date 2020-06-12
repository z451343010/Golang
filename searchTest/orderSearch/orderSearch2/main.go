package main

import "fmt"

// 顺序查找的第二种方式
func main() {

	var names = [5]string{"aaa", "bbb", "ccc", "ddd", "eee"}
	fmt.Println("请输入要查找的名字：")
	var name string
	fmt.Scanln(&name)

	var flag bool

	for i := 0; i < len(names); i++ {
		if name == names[i] {
			fmt.Println("已经找到")
			fmt.Println("下标为：", i)
			flag = true
		}
	}

	if !flag {
		fmt.Println("没有找到")
	}

}
