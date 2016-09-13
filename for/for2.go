package main

import "fmt"

func main() {
	sum := 1
	count := 0
	for ; sum < 1000; {
		sum += sum
		count ++
	}
	fmt.Println(sum, count)
}
