package main

import "fmt"

func main() {
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
