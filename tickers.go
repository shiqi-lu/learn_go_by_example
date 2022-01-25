package main

import (
    "fmt"
    "time"
)

func main() {
    // 打点器和定时器的机制有点相似：使用一个通道来发送数据
    // 这里使用通道内建的 select，等待每 500ms 到达一次的值
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    // 打点器一旦停止，将不能再从它的通道中接收到值
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
