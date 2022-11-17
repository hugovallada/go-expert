package main

import (
	"html/template"
	"log"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

/**
* Template.Must ja junta tudo em um só
 */
func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n"))
	err := t.Execute(os.Stdout, curso)

	if err != nil {
		log.Fatal(err)
	}
}
