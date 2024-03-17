package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Cria um contexto em branco
	ctx := context.Background()

	// Adiciona um timeout ao contexto
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	// Cancela o contexto no fim da execução do programa
	defer cancel()

	// Chama a função que realiza a reserva do hotel usando o contexto criado.
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	// Verifica se o contexto foi cancelado antes de realizar a reserva.
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Context canceled.")
		return
	// Espera 5 segundos para realizar a reserva do hotel.
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked successfully.")
	}
}
