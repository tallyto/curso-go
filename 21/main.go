package main

import (
	"fmt"

	"github.com/tallyto/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())
}
