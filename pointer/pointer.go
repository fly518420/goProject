package main

import "fmt"

func main() {
	i := 42
	
	fmt.Println(i)	
	p := &i
	fmt.Println(*p)
	fmt.Println(p)

	*p = 21
	fmt.Println(i)

	j := 74
	p = &j
	*p = *p /37
	fmt.Println(j)
	fmt.Println(*p)
}
