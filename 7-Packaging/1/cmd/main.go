package main

import (
	"fmt"

	"github.com/fredericobenevides/posgoexpert/7-Packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
	fmt.Println(m.Add())
	// fmt.Println(m.X)
}
