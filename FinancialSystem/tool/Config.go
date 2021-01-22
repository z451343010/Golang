package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	LogPath  string         `json:"log_path"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

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
