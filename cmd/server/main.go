package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	const connection string = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connection)
	if err != nil {
		log.Fatalf("error connecting to amqp: %v", err)
	}
	defer conn.Close()

	fmt.Println("Peril game server connected to RabbitMQ!")

	// wait for ctrl+c
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	_, ok := <-signalCh
	if ok {
		fmt.Println("Shutting down amqp server")
	}
}
