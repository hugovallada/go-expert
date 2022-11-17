package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var Red = "\033[31m"
var Reset = "\033[0m"

func main() {
	log.Println(Red, "Iniciando a chamada http Get...", Reset)
	defer log.Println(Red, "Chamada http Get executada com sucesso", Reset)
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
