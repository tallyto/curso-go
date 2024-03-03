package main

import "fmt"

func main() {
	numeros := []string{"um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez"}

	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	for key, value := range numeros {
		fmt.Println(key+1, value)
	}

	i := 0

	for i < 10 {
		fmt.Println(i + 1)
		i++
	}

	for {
		fmt.Println("loop infinito")
	}
}
