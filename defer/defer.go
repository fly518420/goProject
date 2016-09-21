package main

import "fmt"

func main() {
	defer fmt.Println("我是延迟执行的")

	fmt.Println("后来居上，牛逼不")
}
