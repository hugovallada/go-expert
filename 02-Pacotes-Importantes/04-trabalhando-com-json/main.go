package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"` // se colocar `json:"-"` desconsidera essa variável
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	Jsonify(conta)
	Encoder(conta)

	jsonPuro := []byte(`{"n":2, "s":200}`)
	Structfy(jsonPuro)
}

func Jsonify(conta Conta) { // Marshal vc salva o valor
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func Encoder(conta Conta) { // Encode vc ja manda o valor pra outro luga
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(conta); err != nil {
		panic(err)
	}
}

func Structfy(jsonValue []byte) {
	var conta Conta
	if err := json.Unmarshal(jsonValue, &conta); err != nil {
		panic(err)
	}
	fmt.Println(conta)
}
