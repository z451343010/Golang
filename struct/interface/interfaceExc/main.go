package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
	实现对Hero结构体切片的排序
*/

type Hero struct {
	Name string
	Age  int
}

/*
	Interface接口
	type Interface interface {
		Len() int
		Less(i,j int) bool
		Swap(i,j int)
	}
*/

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 实现sort包的Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	/* 在Golang中，交换可以用别的形式来表示
	 如：
		var i int = 10
		var j int = 20
		i,j = j,i
		这样 i 就会等于 20
		j 就会等于 10
	*/
}

func main() {

	// 测试排序
	var heroes HeroSlice

	for i := 0; i < 10; i++ {

		hero := Hero{
			Name: fmt.Sprintf("Hero %d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}

		// 将 hero append 到 heros 切片
		heroes = append(heroes, hero)

	}

	// 看看排序前的顺序
	fmt.Println("before range")
	for _, v := range heroes {
		fmt.Println(v)
	}

	// 调用 sort.Sort
	/*
		sort包里面有个Sort方法
		可以传入Interface类型的接口作为参数
		func Sort(data Interface) {
			n := data.Len()
			quickSort(data,0,n,maxDepth(n))
		}
	*/
	sort.Sort(heroes)

	fmt.Println("after range")
	for _, v := range heroes {
		fmt.Println(v)
	}

	// var intSlice = []int{0, -1, 10, 7, 90}
	// // 切片是引用类型，所以传入的时候不需要加地址符号
	// sort.Ints(intSlice)

}
