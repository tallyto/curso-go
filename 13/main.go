package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("O cliente %s foi desativado", c.Nome)
}

func main() {
	tallyto := Cliente{
		Nome:  "Tallyto",
		Idade: 27,
		Ativo: true,
	}

	tallyto.Endereco = Endereco{
		Logradouro: "Rua das Flores",
		Numero:     123,
		Cidade:     "São Paulo",
		Estado:     "SP",
	}
	tallyto.Cidade = "Rio de Janeiro"

	tallyto.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", tallyto.Nome, tallyto.Idade, tallyto.Ativo)
	fmt.Println(tallyto.Endereco)
}
