package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.New("content.html")
		// Registra funções que podem ser usadas no template
		tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
		tmp, _ = tmp.ParseFiles(templates...)
		err := tmp.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 30},
			{"Python", 20},
			{"C#", 10},
			{"PHP", 5},
			{"C++", 3},
			{"Ruby", 2},
		})
		if err != nil {
			panic(err)
		}

	})

	http.ListenAndServe(":3001", nil)

}
