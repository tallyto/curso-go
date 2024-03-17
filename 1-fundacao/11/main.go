package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	tallyto := Cliente{
		Nome:  "Tallyto",
		Idade: 27,
		Ativo: true,
	}

	tallyto.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", tallyto.Nome, tallyto.Idade, tallyto.Ativo)
}
