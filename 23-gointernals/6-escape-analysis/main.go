package main

// func add(a, b int) int {
// 	sum := a + b // sum é alocado na stack
// 	return sum   // liberada quando a função termina
// }

// type User struct {
// 	Name string
// }

// func (u *User) PrintName() {
// 	u.Name = "João"
// 	println(u.Name)

// }

// // A variável user é um ponteiro para o objeto User.
// // O ponteiro user é retorno para a função createUser
// // O Objeto User precisa continuar existindo após o retorno da função,
// // pois o ponteiro é usado fora do escopo da função.
// // Portanto, user é alocada na heap.
// func NewUser(name string) *User {
// 	user := &User{Name: name} // user é alocada na heap porque retorna um ponteiro
// 	return user
// }

// Armazenzamento em Estruturas de Dados
// Se uma variável local é armazenada em uma estrutura de dados que sobrevive ao escopo da função,
// ela deve ser alocada na heap.
// func storeInMap() map[string]*int {
// 	m := make(map[string]*int) // m é um mapa que pode sobreviver ao escopo da função
// 	i := 42                    // i é uma variável local
// 	m["key"] = &i              //  o ponteiro para i é armazenada no mapa
// 	return m                   // retornamos o mapa
// 	// IMPORTANTE: se m não fosse retornado, ele não seria alocado para o heap
// }

// func storeInMapWithNoReturn() {
// 	m := make(map[string]*int) // m não escapa para heap
// 	i := 42                    // i é uma variável local
// 	m["key"] = &i              //  o ponteiro para i é armazenada no mapa

// }

// Se uma variável local é usada dentro de uma goroutine, ela deve ser alocada na heap,
// pois a goroutine pode continuar em execução após a função retornar
// func startGoroutine() {
// 	i := 42 // i é uma variável local
// 	go func() {
// 		fmt.Println(i) // i escapa para heap
// 	}()
// }

// função recursiva que poderia gerar um stack overflow
func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

// go build -gcflags="-m" main.go
// go install github.com/go-delve/delve/cmd/dlv@latest
// go build -gcflags="all=-N -l" -o recursive main.go
// dlv debug
// break main.recursive
// continue
// step...step...step...
// stack
func main() {
	recursive(10)
	// println(add(1, 2))
	// fmt.Println(add(1, 2))

	// var p *int
	// {
	// 	i := 42
	// 	p = &i // i escapa para a heap porque p é uma variável de escopo maior
	// }
}
