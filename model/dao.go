package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"userOs/util"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	path   string
	suffix = ".sql"
	Models map[string]interface{}
)

type Model interface {
	ToString() string
	Save() bool
}

/**
初始化：
1。数据配置文件
2。模型
3。用户数据
*/
func init() {
	path = util.GetConfig().DataPath
	Models = make(map[string]interface{}, 0)
	Models["user"] = NewUser

	//init user
	UserDatas = make(map[string]Model, 0)
	RfData("user", "username", UserDatas)

}

/**
读取用户信息
*/
func RfData(name string, primary string, datas map[string]Model) error {
	//打开数据文件
	f, err := os.Open(path + name + suffix)
	if err != nil {
		fmt.Print("读取文件错误1", err)
		return errors.New("打开数据文件失败")
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var field []string
	var dataSlice []string
	for {
		data, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				//fmt.Print("is eof", err)
				break
			}
			fmt.Print("读取文件失败2", err)
			return errors.New("读取数据文件失败")
		}
		dataSlice = strings.Split(strings.TrimSpace(string(data)), ",")
		//fmt.Print(dataSlice)
		if field == nil {
			field = dataSlice
		} else {

			//u := User{}
			//uValue := reflect.ValueOf(u)
			//for i:=0; i<uValue.NumField(); i++ {
			//	fmt.Printf("type:%T\n",uValue.Field(i))
			//}

			toModel(name, primary, datas, field, dataSlice)
		}
	}
	return nil
}

/**
根据字段类型 设置模型
*/
func toModel(name string, primary string, datas map[string]Model, field, data []string) {
	modelV := reflect.ValueOf(Models[name])
	newModel := modelV.Call([]reflect.Value{})[0]
	var primaryKey string
	for i := 0; i < len(data); i++ {
		//获取类型
		modelT := newModel.Elem().FieldByName(field[i]).Type().Name()
		//fmt.Printf("type:%T,model type:%s\n", data[i], modelT)
		//反射调用赋值方法
		modelSet := newModel.MethodByName("Set" + strings.Title(field[i]))
		//根据类型赋值model
		modelSet.Call([]reflect.Value{reflect.ValueOf(toType(data[i], modelT))})

		if field[i] == primary {
			primaryKey = data[i]
		}
	}
	datas[primaryKey] = newModel.Interface().(Model)
}

func toType(data, dtype string) interface{} {
	switch dtype {
	case "int":
		data, err := strconv.Atoi(data)
		if err != nil {
			fmt.Print(err)
			return nil
		}
		return data
	}
	return data
}

/**
写入数据文件
*/
func WfData(name string, datas map[string]Model) bool {
	dataStr := getModelString(datas)
	fmt.Println(dataStr)
	//打开数据文件
	f, err := os.OpenFile(path + name + suffix, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("读取文件错误1", err)
		return false
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	l, err := buf.WriteString(dataStr)
	fmt.Println("l:", l)
	if l<len(dataStr) {
		fmt.Println(err)
	}
	errF := buf.Flush()
	fmt.Println("errF:", errF)
	return true
}
/**
将模型转化为字符串
 */
func getModelString(models map[string]Model) string {
	var field string
	var data string
	for _, model := range models {
		modelV := reflect.ValueOf(model)
		if field == "" {
			for i := 0; i < modelV.Elem().NumField(); i++ {
				field = field + "," + modelV.Elem().Type().Field(i).Name //Type结构体 Filed方法
			}
			field = strings.TrimPrefix(field, ",")
		}
		data = data + model.ToString() + "\n"
	}
	return field + "\n" + data
}
