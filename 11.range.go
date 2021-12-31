package main

import "fmt"

func main() {
    // 使用 range 来对 slice 中的元素求和
    // 数组也可以用这种方法初始化并赋初值
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // range 在数组和 slice 中提供对每项的索引和值的访问
    for i, num := range nums {
        if num == 3 {
            fmt.Println("indenx:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // range 也可以只遍历 map 的键
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // range 在字符串中迭代 unicode 码点(code point)
    // 第一个返回值是字符的起始字节位置，然后第二个是字符本身
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
