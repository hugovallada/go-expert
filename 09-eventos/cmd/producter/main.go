package main

import "github.com/hugovallada/go-expert/fcutils/pkg/queue"

func main() {
	ch, err := queue.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	queue.Publish(ch, "Hello World", "amq.direct")
}
