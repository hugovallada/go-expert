package main

import (
	"exportacao_objetos/math"
	"fmt"
)

func main() {
	mathSt := math.NewMath(5, 5)
	fmt.Println(mathSt)
	fmt.Println(mathSt.Add())
}
