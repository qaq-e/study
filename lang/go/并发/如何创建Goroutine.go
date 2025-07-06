package main
import(
    "fmt"
)
func main(){
    go testgo() //使用关键字go调用函数或者方法 开启一个goroutine
    for i:=0;i<10;i++{
        fmt.Println(i)
    }
    fmt.Println("main 函数结束")
    
}

//自定义函数
func testgo(){
    for i:=0;i<10;i++{
        fmt.Println("测试goroutine",i)
    }
}
