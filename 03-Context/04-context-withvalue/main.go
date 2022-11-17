package main

import (
	"context"
	"fmt"
)

type Metadata struct {
	Token string
}

func main() {
	ctx := context.WithValue(context.Background(), "metadata", Metadata{Token: "ijdjidjajidija"})
	bookHotel(ctx)
}

// Por convenção, as variaveis de contexto vem sempre como primeiro parametro
func bookHotel(ctx context.Context) {
	metadata := ctx.Value("metadata")
	if mt, ok := metadata.(Metadata); ok {
		fmt.Println(mt.Token)
	}

	var m Metadata = Metadata{}

	if me, ok := metadata.(Metadata); ok {
		m = me
	}

	fmt.Println(m)

}
