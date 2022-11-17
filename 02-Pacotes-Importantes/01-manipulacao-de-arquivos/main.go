package main

import (
	"bufio"
	"fmt"
	"os"
)

var filePath string = "./Pacotes-Importantes/01-manipulacao-de-arquivos/arquivos.txt"

func main() {
	criarArquivo()
	leitura()
	leituraParticionada()
	leituraLinhaPorLinha()
	deletarArquivo()
}

func criarArquivo() {
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Hello, World")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
}

func leitura() {
	arquivo, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))
}

func leituraParticionada() {
	arquivo, err := os.Open(filePath)
	defer arquivo.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}

func leituraLinhaPorLinha() {
	arquivo, err := os.Open("arquivos2.txt")
	if err != nil {
		panic(err)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func deletarArquivo() {
	if err := os.Remove("arquivos2.txt"); err != nil {
		panic(err)
	}
}
