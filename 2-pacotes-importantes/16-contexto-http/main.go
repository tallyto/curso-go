package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	// cria um contexto com um timeout de 1 segundo
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	// cria uma requisição HTTP com o contexto criado
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	// envia a requisição HTTP e recebe a resposta do servidor
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	// lê o conteúdo da resposta e imprime na tela o conteúdo da resposta.
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(respBody))
}
