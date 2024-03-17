package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(imc(106, 1.77))
	fmt.Println(imc(106, 0))
}

func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func imc(peso, altura float64) (float64, error) {
	if altura == 0 {
		return 0, fmt.Errorf("Altura inválida")
	}
	return peso / (altura * altura), nil
}
