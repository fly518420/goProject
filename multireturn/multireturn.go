package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return y, x, z
}

func main() {
	a, b, c := swap("go", "hello", "!")
	fmt.Println(a, b, c)
}
