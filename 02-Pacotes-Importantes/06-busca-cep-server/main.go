package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hugovallada/busca-cep-server/cep"
)

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Método HTTP não aceito")
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("O CEP não foi enviado")
		return
	}

	endereco, err := cep.BuscaCep(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("O CEP %s é inválido", cepParam)
		json.NewEncoder(w).Encode(message)
		return
	}

	if *endereco == (cep.Endereco{}) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("O CEP %s é inválido", cepParam))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(endereco)
}

/**
* Quando quiser trabalhar com a variável usar o Marshal,
* Quando só quiser devolver o json, usar o Encoder
 */
