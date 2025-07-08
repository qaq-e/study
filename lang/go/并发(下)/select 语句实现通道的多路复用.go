package main
// select 语句实现通道的多路复用

select {
case <- ch1:
// 如果ch1成功读到数据，则进行该case处理语句。
case ch2 <- 1:
// 如果成功向ch2写入数据，则进行该case处理语句。
default:
// 如果上面都没有成功，则进入default处理流程。
}