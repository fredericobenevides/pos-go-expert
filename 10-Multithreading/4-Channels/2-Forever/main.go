package main

func main() {
	forever := make(chan bool)

	// forever <- true

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever
}
