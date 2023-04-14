package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"github.com/tomschdev/kafka-bench/src/proto/person"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Set up a Kafka writer to write messages to the 'people' topic.
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "people",
		Balancer: &kafka.LeastBytes{},
	})

	// Create 10 random instances of the Person proto message and send them to Kafka.
	for i := 0; i < 10; i++ {
		person := &person.Person{
			Name:   fmt.Sprintf("Person %d", i+1),
			Height: float32(rand.Intn(100) + 150),
			Weight: float32(rand.Intn(100) + 50),
			Gender: randGender(),
			Race:   randRace(),
		}
		message, err := proto.Marshal(person)
		if err != nil {
			panic(err)
		}
		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: message,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Sent message %d: %v\n", i+1, person)
	}

	// Close the Kafka writer to flush any buffered messages.
	writer.Close()
}

func randGender() Gender {
	if rand.Intn(2) == 0 {
		return Gender_FEMALE
	} else {
		return Gender_MALE
	}
}

func randRace() Race {
	races := []Race{
		Race_WHITE,
		Race_BLACK,
		Race_HISPANIC,
		Race_ASIAN,
		Race_OTHER,
	}
	return races[rand.Intn(len(races))]
}

