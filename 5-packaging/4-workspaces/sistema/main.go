package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tallyto/curso-go/5-packaging/4-workspaces/math"
)

func main() {
	m := math.Math{A: 1, B: 2}

	fmt.Println(uuid.New().String())
	fmt.Println(m.Add())
}
