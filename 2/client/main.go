package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Context que expira em 5 segundos
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	// Fazendo um chamada no nosso servidor
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}
