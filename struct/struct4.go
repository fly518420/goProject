package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var (
		p1 = Point{1, 2}
		p2 = Point{X : 1}
		p3 = Point{X : 1, Y : 2}
		p4 = Point{Y : 2}
		p = &Point{1, 2}
	)
	fmt.Println(p1, p2, p3, p4, p)
}
