package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 存款
func (account *Account) Deposite(money float64, pwd string) {

	// 查看输入的密码是否正确
	if pwd != (*account).Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 查看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	(*account).Balance += money
	fmt.Println("存款成功")

}

// 取款
func (account *Account) Withdraw(money float64, pwd string) {

	// 查看输入的密码是否正确
	if pwd != (*account).Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 查看存款金额是否正确
	if money <= 0 || money > (*account).Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	(*account).Balance -= money
	fmt.Println("存款成功")

}

// 查询余额
func (account *Account) Query(pwd string) {

	if pwd != (*account).Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Println("你的账号为：", (*account).AccountNo,
		"你的账号的余额为：", (*account).Balance)

}

func main() {

	account := Account{

		AccountNo: "gs201421122044",
		Pwd:       "z18960",
		Balance:   100.0,
	}

	account.Query("z18960")
	account.Deposite(100, "z18960")
	account.Query("z18960")

}
