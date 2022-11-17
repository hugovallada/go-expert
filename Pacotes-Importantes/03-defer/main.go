package main

import "fmt"

func main() {
	defer fmt.Println("Primeira Linha, mas com defer será a Ultima Linha")
	fmt.Println("Segunda Linha, mas com defer será a Primeira Linha")
	fmt.Println("Terceira Linha, mas com defer será a Segunda Linha")
}
