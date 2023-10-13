package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Criando um Contexto
	ctx := context.Background()
	// Esse contexto tem até 3 segundos para ser executado
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	// Chamando a função de cancelar no final da execução
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// O Select fica aguardando o resultado, quando o resultado
	// chega, ele toma uma decisão
	select {
	// Caso receba um Contexto Pronto (Quando ele for cancelado)...
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	// Caso passe 5 segundos e não foi cancelado...
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
