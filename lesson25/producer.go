package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	if err != nil {
		fmt.Println(err)
		return
	}
	defer p.Close()
	topic := "lesson25"
	message := "Hello from go fullstack cource"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny},
		Value: []byte(message),
	}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	e := <-p.Events()
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("Ошибка доставки")
	} else {
		fmt.Println("Доставлено")
	}
}
