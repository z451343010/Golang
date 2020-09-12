package service

import (
	"customerManage/model"
	"fmt"
	"strconv"
)

type CustomerService struct {
	customers   []model.Customer // Customer切片
	customerNum int              // customer数量

}

// 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {

	// 初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "zjune", "male", 20, "13599",
		"zxx@gmail.com")
	customerService.customers = append(customerService.customers, *customer)

	return customerService

}

// 返回客户切片
func (this *CustomerService) CusList() []model.Customer {
	return this.customers
}

// 显示客户信息
func (this *CustomerService) ShowCusMessage() {

	customers := this.CusList()

	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}

}

// 添加客户
func (this *CustomerService) AddCustomer(customer *(model.Customer)) {

	fmt.Print("请输入客户的姓名：")
	fmt.Scanln(&customer.Name)
	fmt.Print("请输入客户的性别：")
	fmt.Scanln(&customer.Gender)
	fmt.Print("请输入客户的年龄：")
	fmt.Scanln(&customer.Age)
	fmt.Print("请输入客户的电话号码：")
	fmt.Scanln(&customer.Phone)
	fmt.Print("请输入客户的邮箱地址：")
	fmt.Scanln(&customer.Email)

	this.customerNum++
	customer.Id = this.customerNum

	this.customers = append(this.customers, *customer)

}

// 通过id找到相应的客户
// 返回 index 下标
func (this *CustomerService) FindCusById(id int) int {

	var index int = -1

	for key, value := range this.customers {

		if value.Id == id {

			index = key

		}

	}

	return index

}

// 删除客户
func (this *CustomerService) DeleteCus(id int) bool {

	index := this.FindCusById(id)

	if index == -1 {
		return false
	}

	this.customers = append(this.customers[:index],
		this.customers[index+1:]...)

	return true

}

// 修改客户信息
func (this *CustomerService) EditCus(id int) bool {

	var editChoice string  // 修改的选项
	var index int          // 被修改的客户切片的下标
	var editContent string // 修改的内容

	index = this.FindCusById(id)

	fmt.Print("请输入需要修改的选项：")
	fmt.Scanln(&editChoice)
	fmt.Print("请输入修改的内容：")
	fmt.Scanln(&editContent)

	// stirng 转 int 类型
	// 新的年龄
	newAge, err := strconv.Atoi(editContent)

	switch editChoice {
	case "姓名":
		this.customers[index].Name = editContent
	case "性别":
		this.customers[index].Gender = editContent
	case "年龄":
		if err != nil {
			fmt.Println("输入的年龄有误，错误为：", err)
			return false
		}
		this.customers[index].Age = newAge
	case "电话号码":
		this.customers[index].Phone = editContent
	case "邮箱地址":
		this.customers[index].Email = editContent
	default:
		fmt.Println("输入的选项有误")
		return false
	}

	return true

}
