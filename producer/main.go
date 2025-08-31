package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/klef99/wb-school-l0/producer/dto"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	var tm time.Duration

	tm, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		log.Println("Error parsing timeout.")
		tm = time.Second
	}

	log.Println(tm)

	writer := &kafka.Writer{
		Addr:     kafka.TCP(strings.Split(broker, ",")...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	// создаём writer (producer)
	defer writer.Close()

	ticker := time.NewTicker(tm)

	for {
		select {
		case <-ticker.C:
			val, err := json.Marshal(dto.RandomOrder())
			if err != nil {
				log.Fatal(err)
			}

			msg := kafka.Message{
				Value: val,
			}

			err = writer.WriteMessages(context.Background(), msg)
			if err != nil {
				log.Fatalf("failed to write message: %v", err)
			}

			fmt.Println("Message published to topic:", topic)
		}
	}
}
