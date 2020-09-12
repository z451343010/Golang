package main

import (
	"customerManage/model"
	"customerManage/service"
	"fmt"
)

type CustomerView struct {
	choice string // 接收用户输入的选项
	loop   bool   // 是否循环的显示主菜单
	// 增加一个字段 customerService
	customerService *service.CustomerService
}

// 客户列表
func (cusView *CustomerView) customerList() {
	fmt.Println("---------------------客 户 列 表---------------------")
	fmt.Printf("ID\t姓名\t性别\t年龄\t电话号码\t邮箱\t\n")
	cusView.customerService.ShowCusMessage()
}

// 添加客户
func (cusView *CustomerView) addCus() {

	fmt.Println("---------------------添 加 客 户---------------------")
	var customer model.Customer
	cusView.customerService.AddCustomer(&customer)

}

// 删除客户
func (cusView *CustomerView) deleteCus() {

	var deleteChoice int // 输入的删除客户编号
	var isDelete string  // 输入的是否删除选项

	fmt.Println("---------------------删 除 客 户---------------------")
	fmt.Print("请选择删除客户的编号(-1退出)：")
	fmt.Scanln(&deleteChoice)

	if deleteChoice == -1 {
		return
	}

	fmt.Print("确认是否删除（Y/N）：")
	fmt.Scanln(&isDelete)

	if isDelete == "y" || isDelete == "Y" {

		if cusView.customerService.DeleteCus(deleteChoice) {
			fmt.Println("---------------------删 除 成 功---------------------")
		} else {
			fmt.Println("输入的编号有误")
		}

	}

}

// 修改客户
func (cusView *CustomerView) editCus() {

	var editChoice int // 被修改的客户的编号
	fmt.Println("---------------------修 改 客 户---------------------")
	fmt.Print("请选择修改客户的编号(-1退出)：")
	fmt.Scanln(&editChoice)

	if editChoice == -1 {
		return
	}

	if cusView.customerService.EditCus(editChoice) {
		fmt.Println("---------------------修 改 成 功---------------------")
	} else {
		fmt.Println("修改失败，输入有误")
	}

}

// 显示主菜单
func (cusView *CustomerView) MainMenu() {

	cusView.loop = true

	for {
		fmt.Println()
		fmt.Println("---------------------客户信息管理软件---------------------")
		fmt.Println("                     1 添 加 客 户 ")
		fmt.Println("                     2 修 改 客 户 ")
		fmt.Println("                     3 删 除 客 户 ")
		fmt.Println("                     4 客 户 列 表 ")
		fmt.Println("                     5 退       出 ")
		fmt.Print("请选择（1-5）：")
		fmt.Scanln(&(cusView.choice))

		switch cusView.choice {
		case "1":
			cusView.addCus()
		case "2":
			cusView.editCus()
		case "3":
			cusView.deleteCus()
		case "4":
			cusView.customerList()
		case "5":
			cusView.loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if !cusView.loop {
			break
		}

	}

}

func main() {

	var cusView CustomerView
	cusView.customerService = service.NewCustomerService()
	cusView.MainMenu()

}
