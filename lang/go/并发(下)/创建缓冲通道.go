package main

import (
    "fmt"
)

func main() {
    //定义一个缓冲区大小为5的通道
    ch1 := make(chan int, 5)
    ch1 <- 1 //向缓冲区放入数据1 因为缓冲区的大小为5 放入一个1之后 还有四个空的缓冲区  所以还未阻塞
    ch1 <- 2
    ch1 <- 3
    ch1 <- 4
    ch1 <- 5 //此时缓冲区已经满 如果再加入 则会进入阻塞状态
    //继续添加时会造成死锁 因为缓冲区满了 一直没有读取
    ch1 <- 6 //fatal error: all goroutines are asleep - deadlock!
    fmt.Println("main end")
}
