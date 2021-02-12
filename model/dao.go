package model

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	path  = "/shell/study/golang/src/day03/data/"
	suffix  = ".sql"
	Models map[string]interface{}
)

func init()  {
	Models = make(map[string]interface{}, 0)
	Models["user"] = NewUser
}

func RfData(name string, primary string, datas map[string]Model) error {
	f, err := os.Open(path+name+suffix)
	if err != nil {
		fmt.Print("读取文件错误1", err)
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var field  []string
	var dataSlice []string
	for {
		data,err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print("is eof", err)
				break
			}
			fmt.Print("读取文件失败2", err)
			return err
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

func toModel(name string, primary string, datas map[string]Model, field,data []string){
	modelV := reflect.ValueOf(Models[name])
	newModel := modelV.Call([]reflect.Value{})[0]
	var primaryKey string
	for i := 0; i < len(data); i++ {
		fmt.Print("data value:", data[i])
		//获取类型
		modelT := newModel.Elem().FieldByName(field[i]).Type().Name()
		fmt.Printf("type:%T,model type:%s\n", data[i], modelT)
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

func toType(data ,dtype string) interface{} {
	switch dtype {
	case "int":
		data,err := strconv.Atoi(data)
		if err != nil {
			fmt.Print(err)
			return nil
		}
		return data
	}
	return data
}

func WfData()  {
	
}

