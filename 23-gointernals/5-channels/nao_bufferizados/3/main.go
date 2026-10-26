package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case val := <-ch:
		fmt.Println("Received", val)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
