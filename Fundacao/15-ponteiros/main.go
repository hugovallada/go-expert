package main

import "fmt"

func main() {
	a := 10
	fmt.Println(&a) // & pega o endereço da memória onde a variável está apontado

	var ponteiro *int = &a // * indica um ponteiro -> que representa o endereçamento da memória 
	fmt.Println(ponteiro)

	// variável -> ponteiro que tem um endereço na memória -> valor

	*ponteiro = 20 // Muda o valor que está la no endereço
	fmt.Println(a)

	b := &a
	fmt.Println(b)

	// & pega o ponteiro de uma variável
	// * indica q é um ponteiro
	// Quando colocar o * em um ponteiro, trabalha com o valor do ponteiro

	fmt.Println(*b)
	*b = 30
	fmt.Println(*b)
	fmt.Println(a)

	// Porque se usa ponteiros
	/**
	* Para que o valor q esteja trabalhando, não esteja em um escopo muito local
	*/
	
}