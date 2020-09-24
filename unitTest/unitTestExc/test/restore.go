package test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

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
