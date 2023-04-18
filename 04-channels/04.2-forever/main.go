package main

func main() {
	forever := make(chan bool) // canal vazio

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}

		forever <- true // canal sendo preenchido
	}()

	<-forever // canal sendo lido
}
