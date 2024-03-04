package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/pal", handlerPal)
	http.HandleFunc("/cria-pal", handlerCriaPal)

	fmt.Println("Servidor iniciado. Acesse http://localhost:3001")

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %s\n", err)
	}
}

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type PalWorld struct {
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Ataque int    `json:"ataque"`
	Defesa int    `json:"defesa"`
}

func handlerPal(w http.ResponseWriter, r *http.Request) {
	pal := PalWorld{
		Nome:   "Dragão",
		Tipo:   "Fogo",
		Ataque: 100,
		Defesa: 80,
	}

	data, err := json.Marshal(pal)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	pessoa := Pessoa{Nome: "Gabriel", Idade: 28}

	data, err := json.Marshal(pessoa)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
	}
}

func handlerCriaPal(w http.ResponseWriter, r *http.Request) {
	var pal PalWorld

	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(corpo, &pal)
	if err != nil {
		http.Error(w, "Erro ao fazer parse da resposta", http.StatusBadRequest)
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse: %v\n", err)
		return
	}

	_, err = w.Write(corpo)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
	}

	fmt.Printf("Nome: %s, Tipo: %s, Ataque: %d, Defesa: %d\n", pal.Nome, pal.Tipo, pal.Ataque, pal.Defesa)
}
