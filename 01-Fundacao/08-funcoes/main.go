package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(50, 2))
	fmt.Println(calcs(10, 2))

	valor, err := sumWithErr(25, 25)
	if err != nil {
		panic(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, bool) {
	if result := a + b; result > 50 {
		return result, true
	}
	return a + b, false
}

func calcs(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func sumWithErr(a, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
