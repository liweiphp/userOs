package controller

import (
	"fmt"
	"manager/model"
	"manager/util"
)

type UserController struct {

}

func Login() bool {
	fmt.Print("请输入用户名：")
	username := util.CRead()
	fmt.Print("请输入密码：")
	password := util.CRead()
	user := model.GetUser(username)
	if user == nil {
		fmt.Print("暂无此用户")
		return false
	}

	if user.GetPassword() == password {
		//登陆成功
		fmt.Println("登陆成功")
		session = true
		return true
	} else {
		fmt.Print("密码错误")
		return false
	}
}
