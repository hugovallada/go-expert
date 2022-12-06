package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	canal := make(chan string)

	// Thread 2
	go func() {
		canal <- "Olá Mundo!" // Preenchendo o canal
	}()

	// Thread 1
	msg := <-canal
	fmt.Println(msg)
}
