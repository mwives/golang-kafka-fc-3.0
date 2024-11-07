package main

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "golang-kafka:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"golang-kafka"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
