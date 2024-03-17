package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	handlerError(err)

	// tamanho, err := f.WriteString("Escrevendo no arquivo")
	tamanho, err := f.Write([]byte("Escrevendo no arquivo"))

	handlerError(err)

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	// leitura

	arquivo, err := os.ReadFile("arquivo.txt")
	handlerError(err)

	fmt.Println(string(arquivo))

	// leitura de arquivo em lote

	file, err := os.Open("arquivo.txt")

	handlerError(err)

	reader := bufio.NewReader(file)

	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	file.Close()

	err = os.Remove("arquivo.txt")
	handlerError(err)

}

func handlerError(err error) {
	if err != nil {
		panic(err)
	}
}
