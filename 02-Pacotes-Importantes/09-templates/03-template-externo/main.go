package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	newHtmlFile, err := os.Create("template-processado.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = tmp.Execute(newHtmlFile, Cursos{
		{"Go", 120},
		{"Java", 200},
		{"Python", 40},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
