package main

import (
    "fmt"
)

func main() {
    var str string
    fmt.Println("请输入一段文本")
    fmt.Scanf("%s", &str)
    fmt.Println(str)
    fmt.Printf("字符串的长度 ：%d\n", len(str))
    var count int = 0
    for index, ch := range str {
        fmt.Printf("%d %c\n", index, ch)
        if (ch >= 'a' && ch <= 'z') || 
            (ch >= 'A' && ch <= 'Z') {
            count ++
        }
    }
    fmt.Printf("字母的个数 : %d\n", count)
}

