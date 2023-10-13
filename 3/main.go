package main

import (
	"context"
	"fmt"
)

func main() {
	// Context com Key-Value
	// O Contex pode passar dados de uma lado para o outro
	ctx := context.Background()

	// Passando um Token no Context
	// context.WithValue (chave -> valor)
	ctx = context.WithValue(ctx, "token", "xpto")

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Pegando o Value da Key Token
	token := ctx.Value("token")
	fmt.Println(token)
}
