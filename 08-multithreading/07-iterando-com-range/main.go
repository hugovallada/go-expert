package main

import (
	"fmt"
)

// Sincronização entre Threads
func main() {
	ch := make(chan int)
	go publish(ch)
	consumer(ch)
}

func publish(ch chan int) {
	defer close(ch) // fecha o canal, e após isso, ele não vai mais tentar ler o canal
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
