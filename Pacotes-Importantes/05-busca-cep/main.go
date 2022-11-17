package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}
		var endereco Endereco
		err = json.Unmarshal(res, &endereco) // Poderia usar o NewDecoder passando o req.Body
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(endereco)
		file, err := os.Create("./cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", endereco.Cep, endereco.Localidade, endereco.Uf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao preencher o arquivo: %v\n", err)

		}
	}
}
