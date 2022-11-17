package main

import (
	"fmt"
	"fundacao/pacotes-e-modulos/matematica"

	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("O valor da soma é %v\n", soma)
	carro := matematica.Carro{}
	carro.Andar()

	id := uuid.New()
	fmt.Println(id)
}
