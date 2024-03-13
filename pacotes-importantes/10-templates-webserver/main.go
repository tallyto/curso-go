package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.New("template.html")
		tmp, _ = tmp.ParseFiles("template.html")
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
