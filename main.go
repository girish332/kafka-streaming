package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)

type OrderPlacer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewOrderPlacer(producer *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		producer:   producer,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlacer) placeOrder(orderType string, size int) error {
	format := fmt.Sprintf("%s - %d", orderType, size)
	payload := []byte(format)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	},
		op.deliverych,
	)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
	}
	<-op.deliverych
	return nil

}

func main() {
	topic := "HVSE"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "something",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err, p)
		os.Exit(1)
	}

	op := NewOrderPlacer(p, topic)

	for i := 0; i < 1000; i++ {
		if op.placeOrder("market", i+1); err != nil {
			fmt.Printf("Failed to place order: %s\n", err)
		}
		time.Sleep(3 * time.Second)

	}
}
