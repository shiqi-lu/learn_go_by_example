package main

import "os"

func main() {
    // 运行程序将会导致 panic： 输出一个错误消息和协程追踪信息，并以非零的状态退出程序
    panic("a problem")

    _, err := os.Create("/tpm/file")
    if err != nil {
        panic(err)
    }
}
