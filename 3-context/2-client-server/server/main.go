package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3002", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		// Imprime no command line stdout
		log.Println("Request processada com sucesso!")
		// Imprime no browser
		w.Write([]byte("Request processada com sucesso!"))
		break
	case <-ctx.Done():
		// Imprime no command line stdout
		log.Println("Request cancelada pelo client!")
	}
}
