package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("arr == ", arr)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] == %d\n", i, arr[i])
	}
}
