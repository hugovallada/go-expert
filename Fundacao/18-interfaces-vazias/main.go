package main

import "fmt"

type x interface {} // Seria o any

// type x interface {} é igual a type x = interface{} que é igual a any
func Aceite(valor x) {}

func Teste(valor any) {}

func main() {
	var a interface{} = 10
	var b interface{} = "Hello World"

	showType(a)
	showType(b)
}

func showType(t interface{}) {
	fmt.Printf("Tipo: %T\n", t)
}
