package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Ola")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
