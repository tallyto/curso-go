package main

import "fmt"

var (
	b bool    = true
	c int     = 10
	d string  = "Tallyto"
	e float64 = 10.2
)

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	// for i := 0; i < len(arr); i++ {
	// 	println(arr[i])
	// }

	// range -> retorna dois valores: index e a copia do valor nessa posição
	for i, v := range arr {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	println(arr[0])
	println(len(arr))
	println(arr[len(arr)-1])
}
