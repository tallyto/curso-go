package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.google.com"
	req, err := http.Get(url)
	handlerError(err)

	res, err := io.ReadAll(req.Body)

	handlerError(err)

	fmt.Println(string(res))

	req.Body.Close()

}

func handlerError(err error) {
	if err != nil {
		panic(err)
	}
}
