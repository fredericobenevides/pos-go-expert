package main

import (
	"fmt"
)

func main() {
	start := make(chan bool)

	go func() {
		fmt.Println("Goroutine: Initializing...")
		start <- true // Sinaliza que a inicialização está completa
	}()

	<-start // Espera a inicialização
	fmt.Println("Main: Goroutine completed")
}
