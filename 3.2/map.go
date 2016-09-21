package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1:"java", 2:"golang", 3:"python"}
	for key, value := range m {
		fmt.Println("key : ", key, ", value : ", value)
	}
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
        fmt.Println("--------------------")
	for _, key := range keys {
		fmt.Println("key : ", key)
	}
	fmt.Println("--------------------")
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println("key : ", key)
	}
	fmt.Println("--------------------")
	for _, key := range keys {
		fmt.Println("key:", key, ",value:", m[key])
	}
}
