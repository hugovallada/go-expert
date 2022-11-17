package main

import "fmt"

func main() {
	hello := "Hello"

	fmt.Printf("O tipo é %T\n", hello)
	fmt.Printf("O valor de hello é %v\n", hello)

	type ID int

	var x ID

	x = 10

	fmt.Println(x)
	fmt.Printf("O tipo de x é %T\n", x)
}