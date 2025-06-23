package main

import (
	"fmt"
)

func main() {
	Test1(0)//使用int类型做为参数传入到函数
	Test1("ceshi")//使用string类型做为参数传入到函数
}

//空接口
type T interface {
}

//定义一个函数 接收Test接口类型的数据
func Test1(t T) {
	fmt.Println(t)
}
