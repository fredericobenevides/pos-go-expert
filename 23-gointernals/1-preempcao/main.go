package main

import (
	"fmt"
	"runtime"
)

var (
	output [30]string // 3 times, 10 iterations each.
	oi     = 0
)

func main() {
	runtime.GOMAXPROCS(1) // Or set it through environment variable GOMAXPROCS
	chanFinished1 := make(chan bool)
	chanFinished2 := make(chan bool)

	go loop("Goroutine 1", chanFinished1)
	go loop("Goroutine 2", chanFinished2)
	loop("Main", nil)

	<-chanFinished1
	<-chanFinished2

	for _, l := range output {
		fmt.Println(l)
	}
}

func loop(name string, finished chan bool) {
	for i := 0; i < 1_000_000_000; i++ {
		if i%100_000_000 == 0 {
			output[oi] = name
			oi++
		}
	}

	if finished != nil {
		finished <- true
	}
}
