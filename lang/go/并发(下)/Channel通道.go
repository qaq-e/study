package main

import (
    "fmt"
)

func main() {
    var channle chan int
    fmt.Printf("通道的数据类型:%T,通道的值:%v\n", channle, channle) //

    if channle == nil {
        channle = make(chan int)
        fmt.Printf("通过make创建的通道数据类型:%T,通道的值:%v,\n", channle, channle)
        //make创建后 通道的值为 0xc00005c060 也就是一个内存地址
        //所以channel 是一个引用类型的数据
    }
}
