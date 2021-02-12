package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	v := reflect.ValueOf(&a)
	v.Elem().SetInt(50)
	fmt.Printf("type : %T\n", v)
	fmt.Print(a)
}
