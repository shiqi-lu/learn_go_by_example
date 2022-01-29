package main

import (
    "fmt"
    "sort"
)

// 为了在 Go 中使用自定义函数进行排序，需要一个对应的类型
// 在这里创建了一个 byLength 类型，它只是内建类型 []string 的别名
type byLength []string

// 这里为该类型实现了 sort.Interface 接口的 Len、Less 和 Swap 方法
// 这样就可以使用 sort 包的通用 Sort 方法了
// Len 和 Swap 在各个类型中的实现都差不多
// Less 将控制实际的自定义排序逻辑
// 在这个的例子中，按字符串长度递增的顺序来排序
// 所以这里使用了 len(s[i]) 和 len(s[j]) 来实现 Less
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// 一切准备就绪后，我们就可以通过将切片 fruits 强转为 byLength 类型的切片
// 然后对该切片使用 sort.Sort 来实现自定义排序
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
