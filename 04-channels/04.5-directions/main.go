package main

import "fmt"

func main() {
	data := make(chan string)
	go recebe("Hello", data)
	ler(data)

}

// canal somente recebe dados
func recebe(nome string, canal chan<- string) {
	canal <- nome
}

// canal somente envia dados
func ler(data <-chan string) {
	fmt.Println(<-data)
}
