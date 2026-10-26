package main

import (
	"fmt"
	"runtime"
	"time"
)

// Comando:
// GODEBUG=gctrace=1 go run main.go
// GODEBUG=gctrace=1 GOGC=300 go run main.go
// GODEBUG=gctrace=1 GOGC=-1 go run main.go
func main() {
	// Habilitar o tracing do GC
	// debug.SetGCPercent(-1) // Desativa o GC automático
	// Ajustar o percentual do GC (por exemplo, para 300%)
	// debug.SetGCPercent(300)

	// Função para alocar memória
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// Alocando memória para observar o comportamento do GC
	for i := 0; i < 10; i++ {
		allocateMemory(10 * 1024 * 1024) // Aloca 10 MB de memória
		time.Sleep(1 * time.Second)
	}

	// Exibindo o uso de memória
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
