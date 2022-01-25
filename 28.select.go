package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    // 各个通道将在一定时间后接收一个值， 通过这种方式来模拟
    // 并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "two"
    }()

    // 使用 select 关键字来同时等待这两个值， 并打印各自接收到的值
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
