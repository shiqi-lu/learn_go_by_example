package main

import (
    "flag"
    "fmt"
)

func main() {
    // 基本的标记声明仅支持字符串、整数和布尔值选项
    // 这里声明一个默认值为 "foo" 的字符串标志 word 并带有一个简短的描述。 这里的 flag.String 函数返回一个字符串指针（不是一个字符串值）
    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    // 用程序中已有的参数来声明一个标志也是可以的。 注意在标志声明函数中需要使用该参数的指针
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析
    flag.Parse()

    // 输出解析的选项以及后面的位置参数
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
