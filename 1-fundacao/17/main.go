package main

import "fmt"

type Conta struct {
	saldo float64
}

func NewConta(saldo float64) *Conta {
	return &Conta{saldo}
}

func (c *Conta) simular(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := NewConta(100)
	conta.simular(200)
	fmt.Println(conta.saldo)
}
