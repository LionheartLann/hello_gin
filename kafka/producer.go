package main

/*
https://github.com/segmentio/kafka-go
*/
import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// func produce() {
// 	// set up kafka writer using environment variables
// 	topic := "test-topic"
// 	broker := []string{"localhost:9092"}
// 	w := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers:      broker,
// 		Topic:        topic,
// 		Balancer:     &kafka.LeastBytes{},
// 		BatchSize:    1,
// 		BatchTimeout: 10,
// 	})
// 	ctx := context.Background()
// 	// send message
// 	err := w.WriteMessages(
// 		ctx,
// 		kafka.Message{
// 			Topic: "test_producer",
// 			Value: []byte("Hola!"),
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal("could not write message:", err)
// 	}
// 	fmt.Println("Message sent successfully!")
// }

func produce() {
	// to produce messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func consume() {
	// to consume messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func list_topics() {
	conn, err := kafka.Dial("tcp", "localhost:9094")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
func main() {
	// produce()
	// consume()
	list_topics()
}
