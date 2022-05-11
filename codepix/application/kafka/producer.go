package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	return p
}
