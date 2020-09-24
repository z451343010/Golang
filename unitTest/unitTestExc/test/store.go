package test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (monster Monster) Store(filePath string) bool {

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
