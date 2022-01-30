package main

import (
    b64 "encoding/base64"
    "fmt"
)

func main() {
    data := "abc123!?$*&()'-=@~"

    // Go 同时支持标准 base64 以及 URL 兼容 base64。 这是使用标准编码器进行编码的方法
    // 编码器需要一个 []byte，因此我们将 string 转换为该类型
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // 解码可能会返回错误
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // 使用 URL base64 格式进行编解码
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
