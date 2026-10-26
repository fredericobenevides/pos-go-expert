package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string)

	// Thread 2
	go func() {
		canal <- "Olá, mundo!"
	}()

	msg := <-canal
	fmt.Println(msg)
}
