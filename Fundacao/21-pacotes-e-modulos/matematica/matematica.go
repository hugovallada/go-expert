package matematica

import "fmt"

type number interface {
	~int | ~float64
}

func Soma[T number](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	marca string
}

func (c Carro) Andar() {
	fmt.Println("Carro andando!")
}
