package main

import (
	"fmt"
	"time"
)

type Sender chan<- int

type Receiver <-chan int

func main() {
	var ch = make(chan int, 0)
	number := 6
	go func() {
		var sender Sender = ch
		sender <- number
		fmt.Println("Sent")
	}()
	go func() {
		var receiver Receiver = ch
		fmt.Println("Received!", <-receiver)
	}()
	time.Sleep(time.Second)
}
