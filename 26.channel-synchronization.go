package main

import (
	"fmt"
	"time"
)

// 将要在协程中运行这个函数
// done 通道将被用于通知其他协程这个函数已经完成工作
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {
    // 运行一个 worker 协程，并给予用于通知的通道
    done := make(chan bool, 1)
    go worker(done)

    // 程序将一直阻塞，直至收到 worker 使用通道发送的通知
    <-done
}
