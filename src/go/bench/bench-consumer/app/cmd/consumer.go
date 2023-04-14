package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tomschdev/kafka-bench/src/proto/person"
	"google.golang.org/protobuf/proto"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a new Kafka reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "people",
	})

	// Handle OS signals to gracefully shutdown the reader
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)
	defer close(signals)

	// Start consuming messages
	for {
		// Read the next message from the Kafka topic
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error reading message: %v", err)
		}

		// Unmarshal the message into a Person proto
		var person person.Person
		err = proto.Unmarshal(msg.Value, &person)
		if err != nil {
			log.Fatalf("error unmarshalling message: %v", err)
		}

		// Do something with the person data
		fmt.Printf("Received person: %v\n", person)

		// Check for OS signals to shutdown the reader
		select {
		case sig := <-signals:
			log.Printf("Received %v signal, shutting down...\n", sig)
			err := r.Close()
			if err != nil {
				log.Fatalf("error closing Kafka reader: %v", err)
			}
			return
		default:
		}
	}
}
