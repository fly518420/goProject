package main

import (
	"fmt"
	"math"
)

func max(x, n, y float64) float64 {
	if p := math.Pow(x, n); p < y {
		return p
	}	
	return y
}

func main() {
	fmt.Println(max(3, 2, 10),max(3, 3, 10))
}
