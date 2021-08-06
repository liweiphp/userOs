package controller

import (
	"errors"
	"fmt"
	"userOs/util"
	"reflect"
	"strconv"
	"strings"
)

var (
	next        string
	views       map[string][][3]string
	view        string
	controllers map[string]interface{}
)

func Run() {
	next = "index::Welcome"
	fmt.Println("读取配置文件：",util.GetConfig())
	for {
		result := util.Creturn(util.Cfunc(dispatch))
		if result {
			break
		}
	}

}

func dispatch() (bool, error) {
	//============默认方法调用 start================
	//字符串拆分 控制器&方法
	conAct := strings.Split(next, "::")
	fmt.Println("默认控制器和方法", conAct)
	//控制器
	contro, ok := controllers[conAct[0]]
	if !ok {
		return false, errors.New("无此控制器")
	}
	conV := reflect.ValueOf(contro)
	//调用方法
	conV.MethodByName(conAct[1]).Call([]reflect.Value{})
	//=============默认方法调用 end==============

	action, _ := GetView(view)
	for {
		input := util.CRead()
		if input == "x" {
			return true,nil
		}
		order, err := strconv.Atoi(input)
		if order < len(action) && err == nil {
			next = action[order]
			break
		}
		fmt.Println("输入指令错误")
	}

	return false, nil
}
