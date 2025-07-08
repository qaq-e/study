package main

import (
	"fmt"
)
// 还有一种是，一个通道在一个主goroutine协程里同时进行读和写。也会造成死锁。
func main() {
	c := make(chan int)
	c <- 100 //向通道中写入数据
	a := <-c //读取通道中的数据
	fmt.Println(a)
}