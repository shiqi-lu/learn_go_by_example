package main

import "fmt"

// zeroval 有一个 int 型参数，所以使用值传递
// zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival
func zeroval(ival int) {
    ival = 0
}

// zeroptr 使用了 int 指针
// 函数体内的 *iptr 会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值
// 对解引用的指针赋值，会改变这个指针引用的真实地址的值
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // 通过 &i 语法来取得 i 的内存地址，即指向 i 的指针
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // 指针也可以被打印的
    fmt.Println("pointer:", &i)
}
