package main

import (
	"html/template"
	"net/http"
	"strings"
)

/**
* html e text funcionam da msm maneira
* a diferença é que o html se blinda contra alguns tipos de possíveis ataques
 */

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 80},
			{"Python", 50},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8082", nil)
}
