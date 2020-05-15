package main

import "fmt"

// 根据用户输入的星期，输出对应的菜单
func main() {

	var week string
	fmt.Println("请输入星期：")
	fmt.Scanln(&week)

	switch week {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("输入的星期有误")
	}

}
