package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid, err := uuid.Parse("ola")
	if err != nil {
		panic("Not UUID")
	}
	fmt.Println(uuid)
}
