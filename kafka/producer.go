package main

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// set up kafka writer using environment variables
	topic := "test-topic"
	broker := []string{"localhost:9092"}
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      broker,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 10,
	})

	// send message
	err := w.WriteMessages(
		kafka.Message{
			Value: []byte("Hola!"),
		},
	)
	if err != nil {
		log.Fatal("could not write message:", err)
	}

	fmt.Println("Message sent successfully!")
}
