package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

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

	go func() {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "foo",
			"auto.offset.reset": "smallest",
		})
		if err != nil {
			log.Fatalln(err)
		}

		err = consumer.SubscribeTopics([]string{topic}, nil)
		if err != nil {
			log.Fatalln(err)
		}

		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", string(e.Value))

			case *kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)

			}
		}
	}()

	deliverych := make(chan kafka.Event, 10000)
	for {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("FOO")},
			deliverych,
		)

		if err != nil {
			fmt.Printf("Failed to produce message: %s\n", err)
			os.Exit(1)
		}

		<-deliverych
		time.Sleep(3 * time.Second)
	}

}
