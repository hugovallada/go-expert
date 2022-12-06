package main

import (
	"fmt"

	"github.com/hugovallada/go-expert/fcutils/pkg/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := queue.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go queue.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
