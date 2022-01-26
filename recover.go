package main

import "fmt"

func mayPanic() {
    panic("a problem")
}

func main() {
    // 必须在 defer 函数中调用 recover。 当跳出引发 panic 的函数时，
    // defer 会被激活，其中的 recover 会捕获 panic
    defer func() {
        // recover 的返回值是在调用 panic 时抛出的错误
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    // 这行代码不会执行，因为 mayPanic 函数会调用 panic。
    // main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束
    fmt.Println("After mayPanic()")
}
