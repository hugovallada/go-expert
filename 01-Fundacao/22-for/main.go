package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // for mais comun
		fmt.Println(i)
	}

	numeros := []string{"um", "dois", "tres"}

	for index, numero := range numeros {
		fmt.Printf("%v  ->  %v\n", index, numero)
	}

	for { // loop infinito
		break
		// consumir mensagens de fila e etc...
	}

	i := 0

	for i < 10 { // for no tipo while
		fmt.Println(i)
		i++
	}

}
