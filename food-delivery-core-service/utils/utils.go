package utils

import "github.com/IBM/sarama"

func InitProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9093"}, nil)
	if err != nil {
		panic(err)
	}
	return producer
}
