package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(3 * time.Second)
		w.Write([]byte(fmt.Sprintf("Voce é o visitante nº %d\n", number)))
	})
	http.ListenAndServe(":8081", nil)
}
