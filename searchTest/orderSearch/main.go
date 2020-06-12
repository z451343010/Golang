package main

import "fmt"

// 顺序查找
func main() {

	var names [5]string = [5]string{"zzz", "kkk", "mmm",
		"sss", "ddd"}

	fmt.Println("请输入要查找的名字：")
	var name string
	fmt.Scanln(&name)

	for i := 0; i < len(names); i++ {

		if names[i] == name {
			fmt.Println("找到该值：", name, "下标为：", i)
			break
		} else if i == len(names)-1 {
			fmt.Println("没有找到")
		}

	}

}
