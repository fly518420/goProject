package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c float64
    fmt.Scanf("%f%f%f", &a, &b, &c)
    temp := math.Max(a, b)
    max := math.Max(temp, c)
    fmt.Println(max)
}

