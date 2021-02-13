package controller

import (
	"fmt"
)

var session = false

type IndexController struct {
}

func (controller IndexController) Welcome() {
	fmt.Print("欢迎来到信息管理系统\n")

	view = "user_view"
}

func (controller IndexController) Index() {
	fmt.Println("这是首页")
	view = "index_view"
}

func (controller IndexController) List() {
	fmt.Println("这是列表")
	view = "index_view"
}
