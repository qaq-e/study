package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 36

	// fmt.Println(reflect.TypeOf(j)) //通过反射获取变量a的type类型对象
	// fmt.Println(reflect.ValueOf(a))//通过反射获取变量a的value类型对象
	atype := reflect.TypeOf(a) 
	fmt.Println(atype.Name()) //获取变量a的类型对象
	avalue := reflect.ValueOf(a) //获取变量a的值对象
	fmt.Println(avalue.Int()) //获取具体的数值
}
