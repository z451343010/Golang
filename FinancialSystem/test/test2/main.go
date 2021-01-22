package main

import (
	"fmt"
	"gin/FinancialSystem/tool"
)

func main() {

	password := "gyh123456"
	password = tool.EncoderSha256(password)
	fmt.Println(password)

}
