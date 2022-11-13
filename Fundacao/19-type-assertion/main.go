package main

import "fmt"

func main() {
	var minhaVar interface{} = "Hugo Lopes"

	var novaVar interface{} = 10

	fmt.Println(minhaVar.(string)) // .(tipo) afirma pro sistema q é desse tipo

	if res, ok := novaVar.(int); ok {
		fmt.Println(res)
	} else {
		fmt.Println("Não foi possível fazer o casting")
	}
}
