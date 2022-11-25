package main

import (
	"fmt"
	"secao_modulos/math"
)

func main() {
	fmt.Println("Hello World")
	mathSt := math.Math{A: 5, B: 5}
	fmt.Println(mathSt.Add())
}
