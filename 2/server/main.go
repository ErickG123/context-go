package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Pegando o Context que veio da Requisição
	ctx := r.Context()
	log.Println("Request Iniciado")

	defer log.Println("Request Finalizada")

	select {
	// Espera 5 segundos...
	case <-time.After(5 * time.Second):
		// Imprime no command line
		log.Println("Request Processada com Sucesso")
		// Imprime no browser
		w.Write([]byte("Request Processada com Sucesso"))
	// Caso a Requisição seja cancelada...
	case <-ctx.Done():
		// Imprime no command line
		log.Println("Request Cancelada pelo Cliente")
	}
}
