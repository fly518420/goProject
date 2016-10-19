package main

import (
    "fmt"
)

func main() {
    var sum1, sum2 int
    for i := 0; i <= 100; i++ {
        if i % 2 == 0 {
            sum2 += i
        } else {
            sum1 += i
        }
    }
    fmt.Printf("奇数和:%d\n", sum1)
    fmt.Printf("偶数和:%d\n", sum2)
}

