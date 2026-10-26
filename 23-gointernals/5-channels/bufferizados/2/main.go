package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Cria um channel bufferizado com capacidade 2

	ch <- 42 // Envia o primeiro valor
	ch <- 43 // Envia o segundo valor

	fmt.Println(<-ch) // Recebe e imprime o primeiro valor do channel
	fmt.Println(<-ch) // Recebe e imprime o segundo valor do channel
}
