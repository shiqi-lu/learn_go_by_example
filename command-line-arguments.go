package main

import (
    "fmt"
    "os"
)

func main() {
    // os.Args 提供原始命令行参数访问功能
    // 注意，切片中的第一个参数是该程序的路径， 而 os.Args[1:]保存了程序全部的参数
    fmt.Println(os.Args)
    fmt.Println(os.Args[1:])
    fmt.Println(os.Args[3])
}
