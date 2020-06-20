package main

import "fmt"

// 课堂练习：
// 1）使用map[string]map[string]string 的map类型
// 2）key:表示用户名，是唯一的，不可以重复
// 3）如果某个用户名存在，就将其密码改为"888888"
// 如果不存在就增加这个用户信息（包括昵称nickname,和密码pwd）
// 4）编写一个函数 modifyUser(users map[string]map[string]string,
// 	name string) 完成上诉功能

func modifyUser(users map[string]map[string]string) {

	var isExist bool

	for key, value := range users {
		if key == "xanxus" {
			value["pwd"] = "888888"
			isExist = true
		}
	}

	if !isExist {
		newUser := make(map[string]string)
		newUser["name"] = "xanxuskkk"
		newUser["pwd"] = "888888"
		newUser["nickname"] = "kkk"

		users["xanxus"] = newUser

	}

}

func main() {

	users := make(map[string]map[string]string)

	user1 := make(map[string]string)
	user1["name"] = "xanxus"
	user1["pwd"] = "123456"
	user1["nickname"] = "sss"

	user2 := make(map[string]string)
	user2["name"] = "xanxus123"
	user2["pwd"] = "123456"
	user2["nickname"] = "sss123"

	user3 := make(map[string]string)
	user3["name"] = "kyc"
	user3["pwd"] = "123456"
	user3["nickname"] = "sssKyc"

	users["xanxuskkk"] = user1
	users["xanxusddd"] = user2
	users["mmmsss"] = user3

	modifyUser(users)
	fmt.Println(users)

}
