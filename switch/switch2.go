package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today")
		case today + 1:
			fmt.Println("Tomorrow")
		default:
			fmt.Println("Too far away")
	}
}
