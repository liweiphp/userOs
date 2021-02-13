package controller

import "fmt"

func init() {
	initControllers()
	initViews()
}

func initViews() {
	views = make(map[string][][3]string, 0)
	views["user_view"] = [][3]string{
		0: {
			"user",
			"Login",
			"用户登陆",
		},
		1: {
			"user",
			"Register",
			"用户注册",
		},
	}
	views["index_view"] = [][3]string{
		0: {
			"index",
			"Index",
			"首页",
		},
		1:{
			"index",
			"List",
			"列表",
		},
	}

}

func initControllers() {
	controllers = make(map[string]interface{})
	controllers["index"] = &IndexController{}
	controllers["user"] = &UserController{}
}

func GetView(v string) ([]string, []string) {
	flag := views[v]
	var desc  = make([]string, len(flag))
	var action  = make([]string, len(flag))
	for k, f := range flag {
		action[k] = f[0]+"::"+f[1]
		desc[k] = f[2]
	}
	fmt.Println("请选择你的操作")
	for order, text := range desc {
		fmt.Printf("(%d)%s\n", order, text)
	}
	fmt.Println("(x) 退出")
	return action, desc
}
