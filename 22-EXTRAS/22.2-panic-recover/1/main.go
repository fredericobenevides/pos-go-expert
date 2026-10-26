package main

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				println("Recovered from panic1")
			} else if r == "panic2" {
				println("Recovered from panic2")
			}
		}
	}()

	panic2()
}
