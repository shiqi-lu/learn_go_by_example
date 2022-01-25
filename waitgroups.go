package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    // 这个 WaitGroup 用于等待这里启动的所有协程完成
    // 注意：如果 WaitGroup 显式传递到函数中，则应使用指针
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        // 启动几个协程，并为其递增 WaitGroup 的计数器
        wg.Add(1)
        // 避免在每个协程闭包中重复利用相同的 i 值
        // 更多参考：https://go.dev/doc/faq#closures_and_goroutines
        i := i

        // 将 worker 调用包装在一个闭包中，可以确保通知 WaitGroup 此工作线程已完成
        // 这样，worker 线程本身就不必知道其执行中涉及的并发原语
        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    // 阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成
    wg.Wait()
}
