package tool

import (
	"fmt"
	"log"
	"os"
)

func LogRecoder(cfg *Config, path string, err error) {

	file, err := os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 设置日志的输出路径
	log.SetOutput(file)

	if err != nil {
		log.Println("==========" + path + "==========")
		log.Println(err.Error())
	}

	defer file.Close()

}
