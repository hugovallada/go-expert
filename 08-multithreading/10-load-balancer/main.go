package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	numOfWorkers := 100000
	for i := 1; i <= numOfWorkers; i++ {
		go worker(i, data)
	}
	for i := 0; i < 1000000; i++ {
		data <- i
	}
}

// Load Balancer
func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
