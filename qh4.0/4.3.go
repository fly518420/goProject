package main

import (
    "fmt"
)

func main() {
    var ch rune
    fmt.Scanf("%c", &ch)
    switch {
        case ch >= 'a' && ch <= 'z':
            fmt.Printf("%c\n", ch-32)
        case ch >= 'A' && ch <= 'Z':
            fmt.Printf("%c\n", ch+32)
        case ch >= '0' && ch <= '9':
            fmt.Printf("%c\n", ch)
        case ch == ' ':
            fmt.Println("SPACE")
        default:
            fmt.Println("other")        
    }
}

