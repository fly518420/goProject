package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().Unix())
    var arrs [10]int
    for i := 0; i < 10; i++ {
        arrs[i] = rand.Intn(100)
    }
    for k, v := range arrs{
        fmt.Println(k, v)
    }
    for i := 0; i < len(arrs); i++ {
        fmt.Printf("%d ", arrs[i])
    }
    fmt.Println()
    var max, min int = arrs[0], arrs[0]
    for i := 1; i < len(arrs); i++ {
        if max < arrs[i] {
            max = arrs[i]
        }
        if min > arrs[i] {
            min = arrs[i]
        }
    }
    fmt.Println(max, min)
}

