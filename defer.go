package main

import (
    "fmt"
    "os"
)

// 假设我们想要创建一个文件、然后写入数据、最后在程序结束时关闭该文件
// 这里展示了如何通过 defer 来做到这一切
func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

// 关闭文件时，进行错误检查是非常重要的， 即使在 defer 函数中也是如此
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
