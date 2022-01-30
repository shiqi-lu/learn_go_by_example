package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // os.WriteFile 在 go1.16 新增
    //d1 := []byte("hello\ngo\n")
    //err := os.WriteFile("/tmp/dat1", d1, 0644)
    //check(err)

    f, err := os.Create("/tmp/dat2")
    check(err)

    // 打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close
    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    // 调用 Sync 将缓冲区的数据写入硬盘
    f.Sync()

    // bufio 还提供了的带缓冲的 Writer
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    // 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer
    w.Flush()
}
