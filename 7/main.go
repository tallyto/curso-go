package main

import "fmt"

func main() {
	salarios := map[string]int{
		"João":  1200,
		"Maria": 1500,
		"José":  2000,
	}

	delete(salarios, "João")

	salarios["Pedro"] = 2500
	salarios["Kamila"] = 3000

	sal := make(map[string]int)

	salarios["Kamila"] = salarios["Kamila"] * 5

	fmt.Println(sal)

	fmt.Println(salarios)

	for key, salario := range salarios {
		fmt.Printf("%s - %d\n", key, salario)
	}

	for _, value := range salarios {
		fmt.Println(value)
	}
}
