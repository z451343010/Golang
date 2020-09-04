package utils

import "fmt"

type FamilyAccount struct {
	choice      string  // 操作选项
	loopFlag    bool    // 控制是否退出循环的变量
	balance     float64 // 余额
	addMoney    float64 // 收入
	reduceMoney float64 // 支出
	note        string  // 说明
	details     string  // 收支明细
}

// 显示收支明细
func (familyAccount *FamilyAccount) ShowDetails() {

	fmt.Println("---------------当前收支明细记录---------------")
	fmt.Println("收支\t     账户金额\t     收支金额\t     说明\t  ")
	fmt.Println((*familyAccount).details)

}

// 登记收入
func (familyAccount *FamilyAccount) Income() {

	fmt.Println("收入的金额：")
	fmt.Scanln(&((*familyAccount).addMoney))
	fmt.Println("收入的说明：")
	fmt.Scanln(&((*familyAccount).note))
	(*familyAccount).balance += (*familyAccount).addMoney
	// details = "收入     " + string(balance) + "     " + string(addMoney) + "     " + string(note)
	(*familyAccount).details += fmt.Sprintf("收入\t     %v\t     %v\t     %s\t  \n",
		(*familyAccount).balance,
		(*familyAccount).addMoney, (*familyAccount).note)

}

// 登记支出
func (familyAccount *FamilyAccount) Pay() {

	fmt.Println("支出的金额：")
	fmt.Scanln(&((*familyAccount).reduceMoney))
	fmt.Println("支出的说明：")
	fmt.Scanln(&((*familyAccount).note))

	// 判断余额
	if (*familyAccount).reduceMoney > (*familyAccount).balance {
		fmt.Println("余额不足")
		return
	}

	(*familyAccount).balance -= (*familyAccount).reduceMoney

	(*familyAccount).details += fmt.Sprintf("支出\t     %v\t     %v\t     %s\t  \n",
		(*familyAccount).balance, (*familyAccount).reduceMoney,
		(*familyAccount).note)

}

// 编写要给工厂模式的构造方法
// 返回一个 *FamilyAccount 实例
func NewAccount() *FamilyAccount {

	return &FamilyAccount{

		choice:      "",
		loopFlag:    true,
		balance:     10000.0,
		addMoney:    0.0,
		reduceMoney: 0.0,
		note:        "",
		details:     "",
	}

}

// 主界面
func (familyAccount *FamilyAccount) MainMenu() {

	for {

		fmt.Println()
		fmt.Println("---------------家庭收支记账软件---------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退    出")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&((*familyAccount).choice))

		switch (*familyAccount).choice {
		case "1":
			(*familyAccount).ShowDetails()
		case "2":
			(*familyAccount).Income()
		case "3":
			(*familyAccount).Pay()
		case "4":
			(*familyAccount).loopFlag = false
		default:
			fmt.Println("输入的选项有误，请重新输入")

		}

		if (*familyAccount).loopFlag == false {
			fmt.Println("程序已经退出")
			break
		}

	}

}
