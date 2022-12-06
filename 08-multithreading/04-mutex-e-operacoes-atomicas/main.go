package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Voce é o visitante nº %d\n", number)))
	})
	http.ListenAndServe(":8081", nil)
}
