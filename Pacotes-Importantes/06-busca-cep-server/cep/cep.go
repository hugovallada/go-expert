package cep

import (
	"encoding/json"
	"net/http"
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

func BuscaCep(cep string) (*Endereco, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var endereco Endereco

	json.NewDecoder(resp.Body).Decode(&endereco)
	/**
	* Outra opção, seria dar um ioutil.ReadAll(resp.Body) e fazer o unmarshar com o body retornado
	 */

	return &endereco, nil
}
