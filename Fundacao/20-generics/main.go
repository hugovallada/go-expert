package main

import "fmt"

// Constraints
type Number interface {
	~int | float64
}

type MyNumber int // Não entra no caso do int
// Precisa ou declarar o MyNumber, ou antes do int colocar ~, q significa que vai considerar até os typealiases de int

// func Soma[T int | float64](m map[string]T) T {
// 	var soma T
// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Hugo": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Hugo": 2000.00, "João": 3500.00, "Maria": 5000.50}
	m3 := map[string]MyNumber{"Hugo": 1000, "João": 2000, "Maria": 3000}

	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))
	fmt.Println(Soma(m3))

	fmt.Println(Compara("ola", "ola"))
}
