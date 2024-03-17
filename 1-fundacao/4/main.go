package main

// Importação de pacotes
import "fmt"

var (
	b bool    = true
	c int     = 10
	d string  = "Tallyto"
	e float64 = 10.2
)

func main() {
	// Verificar o tipo de uma variável
	fmt.Printf("O tipe de E é %T", e)

	// Verificar o valor de uma variável
	fmt.Printf("O valor de E é %v", e)
}
