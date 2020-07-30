package model

import "fmt"

type account struct {
	AccountNum string
	password   string
	balance    float64
}

// 工厂模式生成 account 结构体
func NewAccount(accountNum string) *account {
	return &account{
		AccountNum: accountNum,
	}
}

// 密码
func (account *account) SetPassword(password string) {

	if len(password) >= 6 && len(password) <= 10 {
		(*account).password = password
	} else {
		fmt.Println("密码长度应在6到10之间")
	}

}

func (account *account) GetPassword() string {

	return (*account).password

}

// 余额
func (account *account) SetBalance(balance float64) {

	if balance >= 20 {
		(*account).balance = balance
	} else {
		fmt.Println("余额应该大于20")
	}

}

func (account *account) GetBalance() float64 {
	return (*account).balance
}
