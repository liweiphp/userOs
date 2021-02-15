package controller

import (
	"fmt"
	"manager/model"
	"manager/util"
	"strconv"
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
	user := model.NewUser()
	fmt.Println("注册用户")
	fmt.Println("请输入用户名：")
	username := util.CRead()
	user.SetUsername(username)
	fmt.Println("请输入密码：")
	password := util.CRead()
	fmt.Println("确认密码：")
	rePassword := util.CRead()
	if password != rePassword {
		fmt.Println("输入密码不一致，请重新输入")
		return false
	}
	user.SetPassword(password)
	fmt.Println("请输入年龄：")
	age := util.CRead()
	age1, _ := strconv.Atoi(age)
	user.SetAge(age1)
	fmt.Println("请输入性别：")
	sex := util.CRead()
	user.SetSex(sex)
	user.Save()
	view = "user_view"
	return true
}
