package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
	CEP    string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Composição
	// Address Endereco -> isso não seria composição, seria só um tipo
}

func main() {
	hugo := Cliente{Nome: "Hugo", Idade: 26, Ativo: true}
	fmt.Println(hugo.Nome)
	fmt.Printf("Cliente: %s tem %d anos e está com status ativo: %v\n", hugo.Nome, hugo.Idade, hugo.Ativo)

	hugo.Ativo = false
	fmt.Println(hugo.Ativo)

	hugo.Cidade = "Ribeirão Preto"
	fmt.Println(hugo)

	endereco := Endereco{
		Rua:    "Espec",
		Numero: 30,
		Cidade: "Ribeirão Preto",
		Estado: "SP",
		CEP:    "109202191",
	}

	hugo.Endereco = endereco

	fmt.Println(hugo)
}
