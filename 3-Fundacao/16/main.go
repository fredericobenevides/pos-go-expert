package main

import "fmt"

type x interface{}

func main() {
	var minhaVar interface{} = "Frederico Benevides"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o result de ok é %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
