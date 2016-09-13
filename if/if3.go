package main

import "fmt"

func max(x, y int) (result int) {
	if x > y {
		result = x
	} else {
		result = y
	}
	return
}

func main() {
	fmt.Println(max(2, 5), max(8, 3))
}
