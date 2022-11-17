package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Timeout: 1 * time.Microsecond,
	}

	resp := try(c.Get("http://google.com"))

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}

func try[T any](resp T, err error) T {
	if err != nil {
		panic(err)
	}
	return resp
}
