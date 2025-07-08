package main

import (
    "fmt"
    "sync"
)

//创建一个同步等待组的对象
var wg sync.WaitGroup

func main() {
    wg.Add(4) //设置同步等待组的数量
    go Sale1()
    go Sale2()
    go Sale3()
    wg.Wait() //主goroutine进入阻塞状态
    fmt.Println("main end...")
}

func Sale1() {
    fmt.Println("func1...")
    wg.Done() //执行完成 同步等待数量减1
}
func Sale2() {
    defer wg.Done()
    fmt.Println("func2...")
}
func Sale3() {
    defer wg.Done() //使用延时执行来减去执行组的数量
    fmt.Println("func3...")
}

