package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	hugo := Cliente{Nome: "Hugo", Idade: 26, Ativo: true}
	fmt.Println(hugo.Nome)
	fmt.Printf("Cliente: %s tem %d anos e está com status ativo: %v\n", hugo.Nome, hugo.Idade, hugo.Ativo)

	hugo.Ativo = false
	fmt.Println(hugo.Ativo)
}
