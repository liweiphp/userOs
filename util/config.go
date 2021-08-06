package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Config struct {
	BasePath string `json:"base_path"`
	DataPath string `json:"data_path"`
}

var (
	instance *Config
	conf = flag.String("conf", "/usr/local/go/src/userOs/etc/config.json", "config path")
)

func init()  {
	//解析命令行
	flag.Parse()
	NewConfig(*conf)
}

func NewConfig(conf string) {
	if instance != nil {
		return
	}
	c,err := ioutil.ReadFile(conf)
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}
	error := json.Unmarshal(c, &instance)
	if error != nil {
		fmt.Println("解析json失败", error)
		return
	}
}

func GetConfig() Config {
	if instance !=nil {
		return *instance
	}
	return Config{}
}
