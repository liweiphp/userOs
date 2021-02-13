package controller

import (
	"fmt"
	"manager/model"
	"manager/util"
)

type UserController struct {

}

func (u *UserController) Login() bool {
	//view = "user_view"
	fmt.Println("请输入用户名：")
	username := util.CRead()
	fmt.Println("请输入密码：")
	password := util.CRead()
	user := model.GetUser(username)
	if user == nil {
		fmt.Print("暂无此用户")
		return false
	}

	if user.GetPassword() == password {
		//登陆成功
		fmt.Println("登陆成功")
		view = "index_view"
		session = true
		return true
	} else {
		fmt.Println("密码错误")
		return false
	}
}

func (u *UserController) Register() bool {
	fmt.Println("注册用户")
	fmt.Println("请输入用户名：")
	//username := util.CRead()
	fmt.Println("请输入密码：")
	//password := util.CRead()
	fmt.Println("请输入年龄：")
	//age := util.CRead()
	fmt.Println("请输入性别：")
	//sex := util.CRead()

	//modelV := reflect.ValueOf()

	view = "user_view"
	return true
}
