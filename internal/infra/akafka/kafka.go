package akafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafakaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "imersao-go",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	kafakaConsumer.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := kafakaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Println("new msg")
			msgChan <- msg
		}
	}

}
