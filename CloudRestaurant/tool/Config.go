package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

var _cfg *Config = nil

// 解析 app.json 配置文件
func ParseConfig(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	// 把 json 字符串解析成 Config 字段
	err = decoder.Decode(&_cfg)
	if err != nil {
		return nil, err
	}

	return _cfg, err

}
