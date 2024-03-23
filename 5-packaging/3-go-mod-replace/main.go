package main

import (
	"fmt"

	"github.com/tallyto/curso-go/5-packaging/3-go-mod-replace/math"
)

func main() {
	m := math.Math{A: 1, B: 2}

	fmt.Println(m.Add())
}
