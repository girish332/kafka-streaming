package producer

//
//import (
//	"encoding/json"
//	"github.com/IBM/sarama"
//	"kafka-streaming/m/models"
//	"log"
//)
//
//type Producer struct {
//	producer sarama.AsyncProducer
//	config   *sarama.Config
//}
//
//func NewProducer() *Producer {
//	config := sarama.NewConfig()
//	config.Producer.Return.Successes = true
//	config.Producer.Return.Errors = true
//	config.Producer.RequiredAcks = sarama.WaitForAll
//	config.Producer.Retry.Max = 5
//	config.Producer.Retry.Backoff = 1000
//	return &Producer{
//		config: config,
//	}
//}
//
//func (p *Producer) ProduceMessage(message models.Message, topic string) {
//	value, err := json.Marshal(message)
//	if err != nil {
//		log.Printf("Error marshalling message: %v", err)
//		return
//	}
//
//	msg := &sarama.ProducerMessage{
//		Topic: topic,
//		Value: sarama.ByteEncoder(value),
//	}
//
//	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
//	if err != nil {
//		log.Printf("Error creating producer: %v", err)
//		return
//	}
//
//	defer func() {
//		if err := producer.Close(); err != nil {
//			log.Printf("Error closing producer: %v", err)
//		}
//	}()
//
//}
