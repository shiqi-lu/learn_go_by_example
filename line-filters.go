package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // 用带缓冲的 scanner 包装无缓冲的 os.Stdin， 这为提供了一种方便的 Scan 方法
    // 将 scanner 前进到下一个令牌（默认为：下一行）
    scanner := bufio.NewScanner(os.Stdin)

    // Text 返回当前的 token，这里指的是输入的下一行
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)
    }

    // 检查 Scan 的错误。 文件结束符（EOF）是可以接受的，它不会被 Scan 当作一个错误
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
