package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "hello"
    flag := strings.Contains(s, "h")
    fmt.Println(flag)
    index := strings.Index(s, "l")
    fmt.Println(index)
    lastIndex := strings.LastIndex(s, "l")
    fmt.Println(lastIndex)
    count := strings.Count(s, "l")
    fmt.Println(count)
    joinStr := strings.Join([]string{"hello", "world"}, "-")
    fmt.Println(joinStr)
    s = "hello,world"
    ss := strings.Split(s, ",")
    fmt.Println(ss)
    s = " ll "
    s = strings.Trim(s, " ")
    fmt.Println(s)
}

