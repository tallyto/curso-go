package main

import (
	"fmt"

	"github.com/tallyto/curso-go/5-packaging/1-introducao/math"
)

func main() {
	m := math.Math{A: 1, B: 2}

	fmt.Println(m.Add())
	fmt.Println("Hello World!")
}
