package main

import "fmt"

func main() {
	var array [3]string // array tem posições fixas
	array[0] = "Ola"
	array[1] = "Mundo"
	array[2] = "!"

	fmt.Println(array[0])

	for _, palavra := range array {
		fmt.Printf("%v ",palavra)
	}

	fmt.Println()
}