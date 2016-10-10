package main

import (
    "fmt"
    "math"
)

func main() {
    var r, d float64
    fmt.Scanf("%f%f", &r, &d)
    fmt.Printf("X轴坐标 : %.2f\n", r * math.Cos(d))
    fmt.Printf("Y轴坐标 : %.3f\n", r * math.Sin(d))
}