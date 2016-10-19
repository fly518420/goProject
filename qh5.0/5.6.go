package main

import (
    "fmt"
    "math"
)

func main() {
    for i := 1; i <= 1000; i++{
        n := math.Pow(float64(i), 2)
        if int(n) % 10 == int(i) {
            fmt.Println(i)
        }
    }
}

