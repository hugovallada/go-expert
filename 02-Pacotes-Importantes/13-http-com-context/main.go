package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
	fmt.Println()

}
