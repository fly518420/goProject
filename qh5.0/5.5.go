package main

import "fmt"

func main() {
    var str string
    fmt.Scanf("%s", &str)
    fmt.Println(str)
    for k, v := range str {
        fmt.Printf("%d-%c\n", k, v)
    }
}

