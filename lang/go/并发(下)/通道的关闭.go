package main

import (
	"fmt"
)

func main() {
	//创建一个通道用来传递数据
	ch1 := make(chan int)

	//通过子协程往通道中放数据
	go func() {
		fmt.Println("======子协程执行======")
		for i := 0; i < 10; i++ {
			ch1 <- i //往通道中放数据
		}
		close(ch1) //关闭通道
	}()

	//主协程通过for循环来获取通道中的所有数据
	for {
		v, ok := <-ch1 //获取通道的状态以及数据
		if !ok {
			fmt.Println("子协程已将通道关闭")
			break
		}
		fmt.Println("获取到的子协程数据为", v)
	}

	fmt.Println("主协程结束")
}

