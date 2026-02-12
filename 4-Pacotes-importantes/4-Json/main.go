package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 1000}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n":2,"s":2000}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Numero, contaX.Saldo)
}
