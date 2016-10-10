package main

import (
    "fmt"
    "math"
)

func main()  {
    fmt.Println("请输入三个float64类型的数")
    var x, y, z float64
    fmt.Scanf("%f%f%f", &x, &y, &z)
    fmt.Println("求和：", x+y+z)
    fmt.Println("平均值：", (x+y+z)/3)
    sumPow := (math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))
    fmt.Println("平方和：", sumPow)
    fmt.Println("平方和开方：", math.Sqrt(sumPow))
}