package main

import "fmt"

// Revisão de slices

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4"}
	fmt.Println(evento[:2])
	fmt.Println(evento[:0])
	fmt.Println(evento[1:])

	evento = append(evento[:0], evento[1:]...)
	fmt.Println(evento)
}
