package main

import (
    "fmt"
)

func main() {
    fmt.Println("请输入华氏温度")
    var f float64
    fmt.Scanf("%f", &f)
    c := 5 * 1.0 / 9 * (f - 32)
    k := 273.16 + c
    fmt.Printf("摄氏温度%f\n", c)
    fmt.Printf("绝对温度%f\n", k)
}

