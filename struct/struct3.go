package main

import "fmt"

type Point struct {
	X int 
	Y int
}

func main() {
	var p Point
	p = Point{1, 2}
	var newP = Point{2, 3}
	fmt.Println(p.X, p.Y)
	fmt.Println(newP.X, newP.Y)
}
