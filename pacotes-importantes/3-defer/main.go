package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")

	handlerError(err)
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	handlerError(err)

	fmt.Println(string(res))
}

func handlerError(err error) {
	if err != nil {
		panic(err)
	}
}
