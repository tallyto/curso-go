package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(conta)

	if err != nil {
		fmt.Println(err)
	}

	jsonPuro := []byte(`{"numero":2, "saldo":200}`)

	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contaX.Saldo)

}
