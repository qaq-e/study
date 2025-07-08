package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("======子协程执行======")
		data := <-ch1 //从通道中读取数据
		fmt.Println("读取到通道中的数据是:", data)
	}()

	ch1 <- 10 //往通道里放数据
	fmt.Println("======主协程结束======")
}

