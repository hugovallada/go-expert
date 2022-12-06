package main

import "fmt"

// variavel chan <- tipo = Canal só vai receber informações (valores serão enviados pra ele) = Receive Only
// variavel <- chan tipo = Canal só vai ser consumido = Send Only
// variavel chan tipo = Faz qualquer coisa
func main() {
	// channel pode receber, enviar ou fazer ambos

	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

// Receive Only
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Send Only
func ler(data <-chan string) {
	fmt.Println(<-data)
}
