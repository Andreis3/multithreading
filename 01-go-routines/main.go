package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")

	// nada aqui
	// sair do programa
	time.Sleep(15 * time.Second) //segurar processo
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}
