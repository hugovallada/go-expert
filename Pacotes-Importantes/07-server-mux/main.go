package main

import (
	"encoding/json"
	"net/http"
)

// Multiplexer = Onde se conectam os handlers
/**
O Go tem o multiplexer padrão, o problema é q vc não tem controle total sobre o q estpa se conectando, e só pode usar um
Você pode criar o seu próprio server mux para ter esse controle, e ter varios servidores disponiveis com as mesmas rotas, mas server mux diferentes
*/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Ola Mundo")
	})
	mux.Handle("/blog", Blog{Title: "Meu Blog"})
	go http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Ola Hugo")
	})

	http.ListenAndServe(":8081", mux2)
}

/**
* Nesse caso precisa de 1 struct pra cada endpoint
* Tem uma maior flexibilidade
 */
type Blog struct {
	Title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
