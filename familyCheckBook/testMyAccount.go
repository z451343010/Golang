package main

import "fmt"

/*
	简单的家庭收支软件
	(面向过程版本)
*/

func main() {

	var choiceStr string // 操作选项

	var loopFlag bool // 控制是否退出循环的变量
	loopFlag = true

	var balance float64 = 10000.0 // 余额
	var addMoney float64          // 收入
	var reduceMoney float64       // 支出
	var note string               // 说明
	var details string            // 收支明细

	for {

		fmt.Println()
		fmt.Println("---------------家庭收支记账软件---------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退    出")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&choiceStr)

		switch choiceStr {
		case "1":
			fmt.Println(4)
			fmt.Println("---------------当前收支明细记录---------------")
			fmt.Println("收支\t     账户金额\t     收支金额\t     说明\t  ")
			fmt.Println(details)
		case "2":
			fmt.Println("收入的金额：")
			fmt.Scanln(&addMoney)
			fmt.Println("收入的说明：")
			fmt.Scanln(&note)
			balance += addMoney
			// details = "收入     " + string(balance) + "     " + string(addMoney) + "     " + string(note)
			details += fmt.Sprintf("收入\t     %v\t     %v\t     %s\t  \n",
				balance, addMoney, note)
		case "3":
			fmt.Println("支出的金额：")
			fmt.Scanln(&reduceMoney)
			fmt.Println("支出的说明：")
			fmt.Scanln(&note)

			// 判断余额
			if reduceMoney > balance {
				fmt.Println("余额不足")
				break
			}

			balance -= reduceMoney
			// details = "收入     " + string(balance) + "     " + string(addMoney) + "     " + string(note)
			details += fmt.Sprintf("支出\t     %v\t     %v\t     %s\t  \n",
				balance, reduceMoney, note)
		case "4":
			loopFlag = false
		default:
			fmt.Println("输入的选项有误，请重新输入")

		}

		if loopFlag == false {
			fmt.Println("程序已经退出")
			break
		}

	}

}
