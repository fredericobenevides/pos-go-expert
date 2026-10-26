package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine: Working...")
		done <- true // Simula que a goroutine terminou seu trabalho
	}()

	<-done // Espera a goroutine terminar
	fmt.Println("Main: Goroutine completed")
}
