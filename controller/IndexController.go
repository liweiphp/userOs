package controller

import (
	"day03/model"
	"day03/util"
	"fmt"
)

func Login()  {
	fmt.Print("请输入用户名：")
	username := util.CRead()
	fmt.Print("请输入密码：")
	password := util.CRead()
	fmt.Print(password)
	var datas map[string]model.Model

	model.RfData("user", username, datas)
	fmt.Print("get data:", datas)
}
