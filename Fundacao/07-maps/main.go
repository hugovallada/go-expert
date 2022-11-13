package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{
		"Hugo":  8000,
		"João":  1000,
		"Maria": 3000,
	}

	fmt.Println(salarios["Hugo"])
	delete(salarios, "Maria")
	fmt.Println(salarios)
	salarios["Wesley"] = 10000
	fmt.Println(salarios)

	valores := make(map[string]int)

	valores["Hugo"] = 8000

	fmt.Println(valores)

	for profissional, salario := range salarios {
		fmt.Printf("Funcionário: %s - Salário: %d\n", profissional, salario)
	}

	empresas := make(map[string]map[string]int)
	duzi := make(map[string]int)
	liv := make(map[string]int)
	empresas["DUZI"] = duzi
	empresas["LIV"] = liv

	empresas["DUZI"]["Hugo"] = 8000
	empresas["DUZI"]["João"] = 3000
	empresas["LIV"]["Ana"] = 10000
	empresas["LIV"]["Filipe"] = 6000
	empresas["DUZI"]["Delmo"] = 3000

	fmt.Println(empresas)

	for empresa, funcionarios := range empresas {
		for nome, salario := range funcionarios {
			fmt.Printf("Funcionário %v da empresa %v, recebe: R$ %d,00\n", nome, empresa, salario)
		}
	}
}
