package main

import (
	"fmt"
	"time"
)

func pingar(c chan string, alternar chan bool) {
	for {
		<-alternar // Aguarda o sinal para enviar o ping
		c <- "ping"
		alternar <- true // Passa o sinal para a próxima função
	}
}

func pongar(c chan string, alternar chan bool) {
	for {
		<-alternar // Aguarda o sinal para enviar o pong
		c <- "pong"
		alternar <- true // Passa o sinal para a próxima função
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Canal para transmitir as mensagens
	c := make(chan string)

	// Canal de controle para alternar entre "ping" e "pong"
	alternar := make(chan bool, 1)
	alternar <- true // Inicializa o controle para o primeiro "ping"

	// Inicia as goroutines
	go pingar(c, alternar)
	go pongar(c, alternar)
	go imprimir(c)

	// Espera o usuário pressionar Enter para terminar
	var entrada string
	fmt.Scanln(&entrada)
}
