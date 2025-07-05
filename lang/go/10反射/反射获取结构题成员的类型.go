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

	fieldnum := typeofmystruct.NumField() //获取结构体成员字段的数量

	fmt.Println("结构体成员字段的数量:", fieldnum) 

	for i := 0; i < fieldnum; i++ {
		fieldname := typeofmystruct.Field(i) //索引对应的字段信息
		fmt.Println(fieldname)
		name, err := typeofmystruct.FieldByName("Name")//根据指定的字符串返回对应的字段信息
		fmt.Println(name, err)
	}

}
