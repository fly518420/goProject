package main

import (
    "fmt"
    "math"
)

func main() {
    var x float64
    fmt.Scanf("%f", &x)
    var y float64
    if x >= 0 {
        y = (math.Sin(x) + math.Cos(x)) / 2
    } else {
        y = (math.Sin(x) - math.Cos(x)) / 2
    }
    fmt.Println(y)
}

