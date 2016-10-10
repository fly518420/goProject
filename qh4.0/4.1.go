package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t)
    fmt.Printf("星期 ：%s\n", t.Weekday().String())
}

