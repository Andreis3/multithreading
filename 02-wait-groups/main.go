package main

import (
	"fmt"
	"sync"
	"time"
)

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) // esperar as 25 operações terminarem
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)
	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}() // Thread anonima

	// nada aqui
	// sair do programa
	waitGroup.Wait()
}

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		waitGroup.Done()
	}
}
