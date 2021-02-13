package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"testing"
)


func TestJson(t *testing.T)  {
	j := `{
		"base_path" : "bpath",
		"data_path" : "dpath"
	}`
	b := []byte(j)
	var c *Config
	err := json.Unmarshal(b, &c)
	fmt.Println(err)
	fmt.Printf("json type:%T",c.BasePath)
	args := []string{
		"-conf=/shell/study/golang/src/manager/etc/config.json",
	}
	flag.CommandLine.Parse(args)
	fmt.Println("config:", *conf)
}

func TestReadConfig(t *testing.T) {
	fmt.Println(GetConfig().BasePath)
}
