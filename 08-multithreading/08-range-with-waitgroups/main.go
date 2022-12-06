package main

import (
	"fmt"
	"sync"
)

// Sincronização entre Threads
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go publish(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}

func publish(ch chan int, wg *sync.WaitGroup) {
	defer close(ch) // fecha o canal, e após isso, ele não vai mais tentar ler o canal
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
