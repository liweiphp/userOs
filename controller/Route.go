package controller

import (
	"day03/util"
	"fmt"
	"strconv"
)

func Run() {
	dispatch()
}

func dispatch() {
	fmt.Print("欢迎来到信息管理系统\n")
	fmt.Print("请选择你的操作\n")
	oper := []string{"登陆", "注册"}
	for key, value := range oper {
		fmt.Printf("%d %s", key, value)
	}
	flag, err := strconv.Atoi(util.CRead())
	if err != nil {
		fmt.Print(err)
		return
	}
	switch flag {
	case 0:
		Login()
	case 1:
		fmt.Print("register")
	}
}
