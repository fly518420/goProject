package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Printf("\nFizzBuzz\n")
        } else if i % 3 == 0 {
            fmt.Printf("\nFizz\n")
        } else if i % 5 == 0 {
            fmt.Printf("\nBuzz\n")
        } else {
            fmt.Printf("%d ", i)
        }
    }
}

