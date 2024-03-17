package main

import "fmt"

type Number interface {
	~int | float64
}

type MyNumber int

func somaInteiro(m map[string]int) int {
	var total int
	for _, v := range m {
		total += v
	}
	return total
}

func somaFloat(m map[string]float64) float64 {
	var total float64
	for _, v := range m {
		total += v
	}
	return total
}

func soma[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{
		"Tallyto": 1000,
		"João":    1500,
		"Maria":   2000}

	salarios := map[string]float64{
		"Tallyto": 1000.0,
		"João":    1500.20,
		"Maria":   2000.35}

	m2 := map[string]MyNumber{
		"Tallyto": 1000,
		"João":    1500,
		"Maria":   2000}

	fmt.Println(somaInteiro(m))

	fmt.Printf("%.2f\n", soma(salarios))

	fmt.Println(soma(m2))

	fmt.Println(Compara("texto", "texto"))

	fmt.Println(Compara(10, 10.0))
}
