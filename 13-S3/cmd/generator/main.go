package main

import (
	"fmt"
	"os"
)

func main() {
	count := 0

	for count < 10000 {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", count))
		if err != nil {
			panic(err)
		}
		f.WriteString("Hello World")
		f.Close()
		count++
	}
}
