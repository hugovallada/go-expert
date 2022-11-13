package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20, 30, 40, 100, 100))
}

func sum(valores ...int) int {
	var sum int

	for _, valor := range valores {
		sum += valor
	}

	return sum
}
