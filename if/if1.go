package main

import "fmt"

func isOld(age int) bool {
	if age > 18 {
		return true	
	}
	return false
}

func main() {
	fmt.Println(isOld(20), isOld(5))
}
