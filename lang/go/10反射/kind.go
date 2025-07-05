package main

import (
	"fmt"
	"reflect"
)

type mystruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func main() {

	typeofmystruct := reflect.TypeOf(mystruct{})

	fmt.Println(typeofmystruct.Name()) //获取反射类型对象  mystruct

	fmt.Println(typeofmystruct.Kind()) //获取反射类型种类  struct

}
