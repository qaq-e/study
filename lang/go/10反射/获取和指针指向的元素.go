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

	typeofmystruct := reflect.TypeOf(&mystruct{})

	fmt.Println(typeofmystruct.Elem().Name()) //获取指针类型指向的元素类型的名称

	fmt.Println(typeofmystruct.Elem().Kind()) //获取指针类型指向的元素类型的种类

}
