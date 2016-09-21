package main

import "fmt"

func main() {
	fmt.Println("start print")
	
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
