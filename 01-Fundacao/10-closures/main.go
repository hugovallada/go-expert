package main

import "fmt"

func main() {
	multiple := func() int {
		return sum(1, 2, 3, 4) * 2
	}()
	fmt.Println(multiple)
}

func sum(valores ...int) int {
	var sum int

	for _, valor := range valores {
		sum += valor
	}

	return sum
}
