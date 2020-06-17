package main

import "fmt"

// map的增删改查
func main() {

	// map的添加
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "深圳"
	cities["no3"] = "杭州"

	// 因为 no1 已经有值
	// 因此这里相当于修改操作
	cities["no1"] = "上海"

	fmt.Println(cities)

	// map的删除
	delete(cities, "no1")
	fmt.Println(cities)

	// map 的查找
	value, isFand := cities["no2"]

	if isFand {
		fmt.Println("找到了该key,值为：", value)
	} else {
		fmt.Println("未扎到该key")
	}

	// 如果一次性要删除所有的key
	// 1）如果我们要删除map的所有key
	// 没有一个专门的方法一次删除，可以遍历一下key，逐个删除。
	// 2）或者 map = make(...),make一个新的
	// 让原来的成为垃圾，被gc回收

	// 方式2：
	cities = make(map[string]string)
	fmt.Println(cities)

}
