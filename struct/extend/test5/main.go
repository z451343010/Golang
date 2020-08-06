package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

func main() {

	// 嵌套匿名结构体后，也可以在创建结构体变量（实例）时
	// 直接指定各个匿名结构体字段的值。
	tv := TV{Goods{"电视机", 4000}, Brand{"海尔", "山东"}}
	fmt.Println(tv)

	tv2 := TV{
		Goods{
			Price: 5000,
			Name:  "荣耀智慧屏",
		},
		Brand{"华为", "深圳"},
	}

	fmt.Println(tv2)

}
