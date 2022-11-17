package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo float64
}

func (c Conta) simular(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) adicionar(valor float64) {
	c.saldo += valor
}

func (c *Cliente) andou() {
	c.nome = "Hugo Lopes"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	hugo := Cliente{"Hugo"}
	hugo.andou()
	fmt.Println(hugo.nome)
	conta := Conta{100.00}
	conta.simular(200.00)
	fmt.Println(conta.saldo)
	conta.adicionar(200.00)
	fmt.Println(conta.saldo)
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}
