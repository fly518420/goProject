package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
    if n < 100 || n > 999 {
        fmt.Println("非法数据")
        return
    }
    var sum float64
    var temp = n
    for {
        x := temp % 10;
        if x == 0 {
            break
        }
        sum += math.Pow(float64(x), 3)
        temp /= 10
    }
    if int(sum) == n {
        fmt.Println("水仙花数")
    } else {
        fmt.Println("----------");
    }
}

