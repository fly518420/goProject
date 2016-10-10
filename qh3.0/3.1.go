package main

import (
    "fmt"
)

func main() {
    fmt.Println("请输入一个三位数")
    var n int
    fmt.Scanf("%d", &n)
    fmt.Printf("输入三位数：%d\n", n)
    var sum int = 0
    x := n % 10
    for x != 0 {
        sum += x
        n = n / 10
        x = n % 10
    }
    fmt.Printf("各位数求和结果：%d\n", sum)
}

