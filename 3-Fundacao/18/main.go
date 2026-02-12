package main

import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro)
	fmt.Println(carro.Andar())

	fmt.Println(uuid.New())
}
