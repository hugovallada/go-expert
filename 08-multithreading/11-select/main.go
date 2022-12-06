package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

// Select funciona como um switch case para canais
func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, fmt.Sprintf("Hello from RabbitMQ %d", i)}
			time.Sleep(time.Second)
			c1 <- msg
		}

	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, fmt.Sprintf("Hello from Kafka %d", i)}
			time.Sleep(time.Second * 2)
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1.Msg)
		case msg2 := <-c2:
			fmt.Println("Received", msg2.Msg)
		}
	}
}

func selectLoopInfinito() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for {
			time.Sleep(time.Second)
			c1 <- 1
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- 2
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}

func selectFor() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}

func selectBasico() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Received", msg1)
	case msg2 := <-c2:
		fmt.Println("Received", msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")
	default:
		fmt.Println("default") // se todo mundo tiver esperando, cai no default
	}
}
