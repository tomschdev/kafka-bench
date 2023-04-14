package main

import (
	"context"
	"fmt"
	"math/rand"

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
			Name:          fmt.Sprintf("Person %d", i+1),
			Age:           int32(rand.Intn(50)),
			Gender:        randGender(),
			Race:          randRace(),
			BiometricData: randBiometricData(),
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

func randGender() person.Gender {
	if rand.Intn(2) == 0 {
		return person.Gender_FEMALE
	} else {
		return person.Gender_MALE
	}
}

func randBiometricData() *person.BiometricData {
	bd := person.BiometricData{
		Height: &person.Height{
			Feet:   int32(rand.Intn(4) + 4),
			Inches: int32(rand.Intn(6) + 6),
		},
		Weight: &person.Weight{
			Pounds: int32(rand.Intn(100) + 100),
		},
	}
	return &bd
}

func randRace() person.Race {
	races := []person.Race{
		person.Race_UNKNOWN_RACE,
		person.Race_WHITE,
		person.Race_BLACK,
		person.Race_ASIAN,
		person.Race_HISPANIC,
		person.Race_NATIVE_AMERICAN,
	}
	return races[rand.Intn(len(races))]
}
