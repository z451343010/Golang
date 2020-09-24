package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

/*
	单元测试综合案例要求
	1) 编写一个Monster结构体，字段Name,Age,Skill
	2) 给Monster绑定方法Store，可以将一个Monster变量（对象）
	序列化后保存到文件中
	3)给Monster绑定方法ReStore,可以将一个序列化的Monster
	从文件中读取，并反序列化为Monster对象
	4)编写测试用例文件 store_test.go,编写测试用例函数TestStore
	和TestRestore进行测试
*/

type Monster struct {
	Name  string
	Age   int
	Skill string
}

/*
	序列化
	写操作
*/
func (monster *Monster) Store(filePath string) bool {

	// 序列化 monster
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("serialize err = ", err)
	}

	// 序列化后的JSON字符串
	serStr := string(data)

	// 将 JSON 写入文件
	file, wErr := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if wErr != nil {
		fmt.Println("open file err = ", err)
		return false
	}

	// 准备写入文件
	writer := bufio.NewWriter(file)
	writer.WriteString(serStr)
	writer.Flush()

	defer file.Close()

	return true

}

/*
	反序列化
	读操作
*/
func (monster *Monster) ReStore(filePath string) bool {

	file, readErr := os.OpenFile(filePath, os.O_RDONLY, 0666)

	if readErr != nil {
		fmt.Println("open file readErr = ", readErr)
		return false
	}

	defer file.Close()

	var jsonStr string // 接收文件中的json字符串

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		jsonStr += str

		if err == io.EOF {
			break
		}

	}

	fmt.Println(jsonStr)

	err := json.Unmarshal([]byte(jsonStr), monster)
	if err != nil {
		fmt.Println("unserialize err = ", err)
		return false
	}

	fmt.Println("反序列化后的对象为：")
	fmt.Println(*monster)

	return true

}

func main() {

	filePath := "F:\\Golang\\goProgramProject\\src\\unitTest\\unitTestExc\\monsterJson.txt"
	monster := &Monster{
		Name:  "zjune",
		Age:   24,
		Skill: "programing",
	}

	monster.Store(filePath)

	var monster2 Monster

	monster2.ReStore(filePath)

}
