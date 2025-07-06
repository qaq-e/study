package main

import (
	"fmt"
	"time"
)

func main() {
    go testgo1()
    go testgo2()
    for i := 0; i <= 5; i++ {
        fmt.Println("main函数执行", i)
    }
	time.Sleep(3000 * time.Millisecond) // 等待子goroutine执行完毕
	// 注意：在实际应用中，使用 sync.WaitGroup 或 channel 来更优雅地管理 goroutine 的生命周期
	// 这里使用 Sleep 只是为了演示，避免主函数过早结束导致子 goroutine 无法执行完毕
	// 这不是一个好的实践，因为它会阻塞主线程，实际应用中应该使用更好的同步方式
	// 例如使用 sync.WaitGroup 来等待所有 goroutine 完成
	// 或者使用 channel 来通知主 goroutine 子 goroutine 的完成状态
	time.Sleep(1000 * time.Millisecond) // 等待子 goroutine 执行完毕
	// 结束主函数
    fmt.Println("main 函数结束")

}

func testgo1() {
    for i := 0; i <= 10; i++ {
        fmt.Println("测试子goroutine1", i)
    }
}

func testgo2() {
    for i := 0; i <= 10; i++ {
        fmt.Println("测试子goroutine2", i)
    }
}
