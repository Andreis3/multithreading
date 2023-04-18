package main

import "fmt"

// Thread 01
func main() {
	channel := make(chan string) // canal vazio

	// Thread 02
	go func() {
		channel <- "Hello, World!" // canal sendo preenchido
	}()

	// Thread 01
	msg := <-channel // canal sendo lido
	fmt.Println(msg)
}
