package main

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"

	println(<-ch)
	println(<-ch)
}
