package databus

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type TelemetryMessage struct {
	DeviceID   string            `json:"deviceID"`
	Timestamp  string            `json:"timestamp"`
	DeviceType string            `json:"deviceType"`
	Data       map[string]string `json:"data"`
}

func main() {
	broker := "localhost:9092"
	topic := "your_topic_name"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("could not read message %v", err)
		}

		var telemetryMessage TelemetryMessage
		if err := json.Unmarshal(msg.Value, &telemetryMessage); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		log.Printf("received message: %+v", telemetryMessage)
	}
}
