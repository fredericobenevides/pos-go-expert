package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	numWorkers := 3

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			fmt.Printf("Worker %d: Working...\n", id)
			done <- true // Sinaliza que a goroutine terminou o trabalho
		}(i)
	}

	for i := 0; i < numWorkers; i++ {
		<-done // Espera cada goroutine terminar
	}

	fmt.Println("All workers completed.")
}
