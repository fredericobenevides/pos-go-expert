package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// Contador para manter o número de goroutines vazadas
var leakedGoroutines int32

func leakGoroutine() {
	atomic.AddInt32(&leakedGoroutines, 1)
	for {
		// Loop infinito representando o vazamento
	}
}

func leakHandler(w http.ResponseWriter, r *http.Request) {
	go leakGoroutine() // Inicia uma nova goroutine que vai vazar
	fmt.Fprintf(w, "Leaked a goroutine!\n")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Retorna o número atual de goroutines vazadas
	fmt.Fprintf(w, "Current number of leaked goroutines: %d\n", atomic.LoadInt32(&leakedGoroutines))
}

func main() {
	http.HandleFunc("/leak", leakHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server started on :6666")
	if err := http.ListenAndServe(":6666", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

// GOMAXPROCES=1 GODEBUG=schedtrace=1
// go-wrk -c 20 -d 100000 http://localhost:6666/leak

// SCHED 17658ms: gomaxprocs=1 idleproces=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=1 runqueue=382 [55]
// Fila de Execução Global (runqueue=382)
// A file de execução global é uma estrutura de dados mantida pelo scheduler do Go que armazena as goroutines prontas para serem executadas.
